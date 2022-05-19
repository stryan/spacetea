package main

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type placeModel struct {
	list list.Model
}

func newPlaceModel(entries []string, m tea.Model) placeModel {
	var p placeModel
	items := []list.Item{}
	for _, v := range entries {
		items = append(items, item{v, v})
	}
	//w,h
	p.list = list.New(items, list.NewDefaultDelegate(), 0, 0)
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
		p.list.SetWidth(msg.Width)
		return p, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl-c":
			return p, tea.Quit
		case "esc":
			return initialModel(), nil
		case "enter":
			cur := p.list.SelectedItem()
			log.Println(cur.FilterValue())
			return initialModel(), nil
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
