package main

import (
	sim "git.saintnet.tech/stryan/spacetea/simulator"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type craftModel struct {
	list list.Model
}

type craftMsg string

func (c craftModel) buildCraftMsg() tea.Msg {
	i := c.list.SelectedItem().(item)
	return craftMsg(i.ID())
}

func newCraftModel(entries []sim.ItemEntry, m tea.Model) craftModel {
	var c craftModel
	items := []list.Item{}
	for _, v := range entries {
		items = append(items, item{v.Name(), "no description", v.ID()})
	}
	//w,h
	c.list = list.New(items, list.NewDefaultDelegate(), 32, 32)
	c.list.Title = "What do you want to craft?"
	c.list.DisableQuitKeybindings()
	return c
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (c craftModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (c craftModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		c.list.SetSize(msg.Width-h, msg.Height-v)
		return c, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl-c":
			return c, tea.Quit
		case "esc":
			return initMainscreen(), heartbeat()
		case "enter":
			return initMainscreen(), tea.Batch(c.buildCraftMsg, heartbeat())
		}
	}

	var cmd tea.Cmd
	c.list, cmd = c.list.Update(msg)
	return c, cmd
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (c craftModel) View() string {
	return c.list.View()
}
