package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	sim "git.saintnet.tech/stryan/spacetea/simulator"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	Width(24).
	Align(lipgloss.Center).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("63"))

type model struct {
	s     *sim.Simulator
	input textinput.Model
}
type beat struct{}

var simulator *sim.Simulator

func heartbeat() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return beat{}
	})
}
func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "input command"
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		s:     simulator,
		input: ti,
	}
}
func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, heartbeat())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			m.s.Stop()
			return m, tea.Quit
		case tea.KeyEnter:
			m.s.Input(m.input.Value())
			m.input.Reset()
			return m, nil
		case tea.KeyRight:
			m.s.Input("right")
			return m, nil
		case tea.KeyLeft:
			m.s.Input("left")
			return m, nil
		case tea.KeyUp:
			m.s.Input("up")
			return m, nil
		case tea.KeyDown:
			m.s.Input("down")
			return m, nil
		case tea.KeyCtrlF:
			if m.input.Focused() {
				m.input.Blur()
			} else {
				m.input.Focus()
			}
			return m, nil
		case tea.KeyRunes:
			if !m.input.Focused() {
				switch msg.String() {
				case "g":
					m.s.Input("get")
				case "p":
					var res []string
					for k, _ := range m.s.Player.Resources {
						res = append(res, strconv.Itoa(k))
					}
					return newPlaceModel(res, nil), nil
				}
				return m, nil
			}
		}
	case beat:
		//heartbeat
		return m, heartbeat()
	}
	m.input, cmd = m.input.Update(msg)
	return m, cmd

}

func (m model) View() string {
	var render string
	display := lipgloss.JoinHorizontal(0, style.Render(m.s.Place.String()), style.Render(fmt.Sprintf("Player\n%v", m.s.Player.String())))
	render = fmt.Sprintf("%v\n%v\n%v\n", style.Render(fmt.Sprintf("Current Time: %v", strconv.Itoa(m.s.Time))), display, style.Render(m.input.View()))
	return render
}

func main() {
	simulator = sim.NewSimulator()
	simulator.Start()

	if err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Start(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
