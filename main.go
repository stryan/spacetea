package main

import (
	"fmt"
	"os"

	sim "git.saintnet.tech/stryan/spacetea/simulator"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	simulator := sim.NewSimulator()
	simulator.Start()
	parent := parent{initMainscreen(simulator)}
	if err := tea.NewProgram(parent, tea.WithAltScreen()).Start(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
