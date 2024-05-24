package main

import tea "github.com/charmbracelet/bubbletea"

type WindowShowMsg struct {
	msg string
}

func (this WindowShowMsg) update(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.dialog = nil
		}
	}
	return m, nil
}

func (this WindowShowMsg) view(m model) string {
	return this.msg + "\n\n" + SUBTLE_STYLE.Render("press press enter to get back.")
}
