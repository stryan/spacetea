package main

import (
	sim "git.saintnet.tech/stryan/spacetea/simulator"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc, id string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) ID() string          { return i.id }
func (i item) FilterValue() string { return i.title }

type placeModel struct {
	list list.Model
}

type placeMsg string

func (p placeModel) buildPlaceMsg() tea.Msg {
	i := p.list.SelectedItem().(item)
	return placeMsg(i.ID())
}

func newPlaceModel(entries []sim.ItemEntry, m tea.Model) placeModel {
	var p placeModel
	items := []list.Item{}
	for _, v := range entries {
		items = append(items, item{v.Name(), "no description", v.ID()})
	}
	//w,h
	p.list = list.New(items, list.NewDefaultDelegate(), 32, 32)
	p.list.Title = "What do you want to place?"
	p.list.DisableQuitKeybindings()
	return p
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (p placeModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (p placeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		p.list.SetSize(msg.Width-h, msg.Height-v)
		return p, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl-c":
			return p, tea.Quit
		case "esc":
			return initMainscreen(), heartbeat()
		case "enter":
			return initMainscreen(), tea.Batch(p.buildPlaceMsg, heartbeat())
		}
	}

	var cmd tea.Cmd
	p.list, cmd = p.list.Update(msg)
	return p, cmd
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (p placeModel) View() string {
	return p.list.View()
}
