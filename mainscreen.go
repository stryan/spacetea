package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	sim "git.saintnet.tech/stryan/spacetea/simulator"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	Width(32).
	Align(lipgloss.Center).
	Foreground(lipgloss.Color("#F0FFFF")).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("63"))

type model struct {
	s        *sim.Simulator
	input    textinput.Model
	help     help.Model
	keys     keyMap
	next     string
	lastSize tea.WindowSizeMsg
}

type beat struct{}
type simCmd string

func heartbeat() tea.Cmd {
	return tea.Tick(time.Second*1, func(time.Time) tea.Msg {
		return beat{}
	})
}

func (m model) simCommand() tea.Msg {
	if m.next != "" {
		tmp := m.next
		m.next = ""
		return simCmd(tmp)
	}
	return nil
}
func initMainscreen() model {
	ti := textinput.New()
	ti.Placeholder = "input command"
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		s:     simulator,
		input: ti,
		help:  help.New(),
		keys:  keys,
	}
}
func (m model) Init() tea.Cmd {
	return heartbeat()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case beat:
		//heartbeat
		return m, heartbeat()
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
		m.lastSize = msg

	case placeMsg:
		simc := fmt.Sprintf("place %v", string(msg))
		m.s.Input(simc)
		return m, nil
	case craftMsg:
		simc := fmt.Sprintf("craft %v", string(msg))
		m.s.Input(simc)
		return m, nil
	case readMsg:
		pid, err := strconv.Atoi(string(msg))
		if err != nil {
			panic(err)
		}
		return newJView(sim.GlobalPages[pid].Title, sim.GlobalPages[pid].Content), m.GetSize
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		case key.Matches(msg, m.keys.Quit):
			m.s.Stop()
			return m, tea.Quit
		case key.Matches(msg, m.keys.Right):
			m.s.Input("right")
			return m, nil
		case key.Matches(msg, m.keys.Left):
			m.s.Input("left")
			return m, nil
		case key.Matches(msg, m.keys.Up):
			m.s.Input("up")
			return m, nil
		case key.Matches(msg, m.keys.Down):
			m.s.Input("down")
			return m, nil
		case key.Matches(msg, m.keys.Gather):
			m.s.Input("get")
		case key.Matches(msg, m.keys.Place):
			var res entrylist
			for _, k := range sim.GlobalItems {
				if m.s.Player.Resources[k.ID()] != 0 {
					res = append(res, k.(sim.ItemEntry))
				}
			}
			sort.Sort(res)
			return newMenuModel(res, placeMenu), m.GetSize
		case key.Matches(msg, m.keys.Pickup):
			m.s.Input("pickup")
		case key.Matches(msg, m.keys.Destroy):
			m.s.Input("destroy")
		case key.Matches(msg, m.keys.Craft):
			var res entrylist
			for k := range m.s.Player.Craftables {
				res = append(res, sim.GlobalItems[k].(sim.ItemEntry))
			}
			sort.Sort(res)
			return newMenuModel(res, craftMenu), m.GetSize
		case key.Matches(msg, m.keys.Journal):
			var res pagelist
			for k := range m.s.Player.Pages {
				res = append(res, sim.GlobalPages[k])
			}
			sort.Sort(res)
			return newJMenuModel(res), m.GetSize
		}
	}
	m.input, cmd = m.input.Update(msg)
	return m, cmd

}

func (m model) View() string {
	var render string
	if m.help.ShowAll != true {
		playerAndLog := lipgloss.JoinVertical(0, style.Render(fmt.Sprintf("%v", m.s.Player.String())), style.Render(fmt.Sprintf("%v", m.s.Player.Log())))
		placeAndInput := lipgloss.JoinVertical(0, style.Render(m.s.Place.String()), style.Render(m.help.View(m.keys)))
		display := lipgloss.JoinHorizontal(0, placeAndInput, playerAndLog)
		render = fmt.Sprintf("%v\n%v\n", style.Render(fmt.Sprintf("Current Time: %v", strconv.Itoa(m.s.Time))), display)
	} else {
		render = style.Render(m.help.View(m.keys))
	}
	return render
}

func (m model) GetSize() tea.Msg {
	return m.lastSize
}
