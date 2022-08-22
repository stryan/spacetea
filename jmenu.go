package main

import (
	sim "git.saintnet.tech/stryan/spacetea/simulator"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type jEntry struct {
	title, id string
}

func (i jEntry) Title() string       { return i.title }
func (i jEntry) Description() string { return "" }
func (i jEntry) ID() string          { return i.id }
func (i jEntry) FilterValue() string { return i.title }

type jMenuModel struct {
	acc      *Account
	list     list.Model
	lastSize tea.WindowSizeMsg
}

type readMsg string

func newJMenuModel(pages []sim.JournalPage, acc *Account) jMenuModel {
	var jm jMenuModel
	items := []list.Item{}
	for _, v := range pages {
		items = append(items, jEntry{v.Title, v.ID()})
	}
	jm.list = list.New(items, list.NewDefaultDelegate(), 80, 32)
	jm.list.DisableQuitKeybindings()
	jm.list.Title = "Read which entry?"
	jm.acc = acc
	return jm
}

func (j jMenuModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (j jMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		j.list.SetSize(msg.Width-h, msg.Height-v)
		j.lastSize = msg
		return j, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl-c":
			return j, tea.Quit
		case "esc":
			return initMainscreen(j.acc), tea.Batch(j.GetSize, heartbeat())
		case "enter":
			return initMainscreen(j.acc), tea.Batch(j.GetSize, j.buildJmenuMsg, heartbeat())
		}
	}

	var cmd tea.Cmd
	j.list, cmd = j.list.Update(msg)
	return j, cmd

}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (j jMenuModel) View() string {
	return j.list.View()
}

func (j jMenuModel) buildJmenuMsg() tea.Msg {
	if j.list.SelectedItem() == nil {
		return ""
	}
	i := j.list.SelectedItem().(jEntry)
	return readMsg(i.ID())

}

func (m jMenuModel) GetSize() tea.Msg {
	return m.lastSize
}
