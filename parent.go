package main

import tea "github.com/charmbracelet/bubbletea"

type parent struct {
	current tea.Model
}

func (m parent) Init() tea.Cmd {
	return tea.Batch(m.current.Init(), tea.EnterAltScreen)
}

func (m parent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.current.Update(msg)
}

func (m parent) View() string {
	return m.current.View()
}
