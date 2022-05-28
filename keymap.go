package main

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Up      key.Binding
	Down    key.Binding
	Left    key.Binding
	Right   key.Binding
	Help    key.Binding
	Quit    key.Binding
	Gather  key.Binding
	Pickup  key.Binding
	Destroy key.Binding
	Place   key.Binding
	Craft   key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right, k.Gather, k.Place, k.Craft, k.Destroy, k.Pickup, k.Help, k.Quit}, //second column
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("esc", "ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
	Gather: key.NewBinding(
		key.WithKeys("g"),
		key.WithHelp("g", "gather tile resource"),
	),
	Pickup: key.NewBinding(
		key.WithKeys(","),
		key.WithHelp(",", "pickup object"),
	),
	Destroy: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "destroy tile object "),
	),
	Place: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "place object"),
	),
	Craft: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "craft object"),
	),
}
