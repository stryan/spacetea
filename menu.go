package main

import (
	sim "git.saintnet.tech/stryan/spacetea/simulator"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type menutype int

const (
	placeMenu menutype = iota
	craftMenu
)

type item struct {
	title, desc, id string
	entry           sim.ItemEntry
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) ID() string          { return i.id }
func (i item) FilterValue() string { return i.title }

type menuModel struct {
	list list.Model
	kind menutype
}

type placeMsg string
type craftMsg string

func newMenuModel(entries []sim.ItemEntry, i menutype) menuModel {
	var p menuModel
	p.kind = i
	items := []list.Item{}
	for _, v := range entries {
		items = append(items, item{v.Describe(), v.Description(), v.ID().String(), v})
	}
	//w,h
	p.list = list.New(items, list.NewDefaultDelegate(), 80, 32)
	switch i {
	case placeMenu:
		p.list.Title = "What do you want to place?"
	case craftMenu:
		p.list.Title = "What do you want to craft?"
	}
	p.list.DisableQuitKeybindings()
	return p
}

func (p menuModel) buildMenuMsg() tea.Msg {
	if p.list.SelectedItem() == nil {
		return ""
	}
	i := p.list.SelectedItem().(item)
	if p.kind == placeMenu {
		return placeMsg(i.entry.String())
	}
	if p.kind == craftMenu {
		return craftMsg(i.entry.String())
	}
	return ""
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (p menuModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (p menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			return initMainscreen(), tea.Batch(p.buildMenuMsg, heartbeat())
		}
	}

	var cmd tea.Cmd
	p.list, cmd = p.list.Update(msg)
	return p, cmd
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (p menuModel) View() string {
	return p.list.View()
}
