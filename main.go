package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choice int
	dialog Window
	window Window
	time   Time
	data   Data
	status Status
	events []Event
}

func main() {
	initial_model := model{}
	initial_model.window = new(WindowAction)
	p := tea.NewProgram(initial_model)
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}
