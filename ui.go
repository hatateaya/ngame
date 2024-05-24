package main

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	SUBTLE_STYLE   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	CHECKBOX_STYLE = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	MAIN_STYLE     = lipgloss.NewStyle().MarginLeft(2)
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.window = new(WindowQuitting)
			return m, tea.Quit
		}
	}
	return updateWindow(m, msg)
}

func (m model) View() string {
	return MAIN_STYLE.Render("当前时间：" +
		strconv.Itoa(m.time.day) + "日" +
		strconv.Itoa(m.time.hour) + "时\n\n" +
		viewWindow(m) + "\n" +
		SUBTLE_STYLE.Render("press key q, esc or ctrl+c to exit."))
}
