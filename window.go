package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Window interface {
	update(m model, msg tea.Msg) (tea.Model, tea.Cmd)
	view(m model) string
}

func updateWindow(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.dialog != nil {
		return m.dialog.update(m, msg)
	} else {
		return m.window.update(m, msg)
	}
}

func viewWindow(m model) string {
	if m.dialog != nil {
		return m.dialog.view(m)
	} else {
		return m.window.view(m)
	}
}

type WindowQuitting struct {
}

func (this WindowQuitting) update(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (this WindowQuitting) view(m model) string {
	return "感谢您的游玩！"
}
