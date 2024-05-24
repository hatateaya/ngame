package main

import tea "github.com/charmbracelet/bubbletea"

type WindowSelector struct {
	options []WindowSelectorOption
}

func (this WindowSelector) update(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			if m.choice > 0 {
				m.choice--
			}
		case tea.KeyDown:
			if m.choice < len(this.options)-1 {
				m.choice++
			}
		case tea.KeyEnter:
			m = this.options[m.choice].on(m)
			m.choice = 0
		}
	}
	return m, nil
}

func (this WindowSelector) view(m model) string {
	var s string = ""
	for i, option := range this.options {
		s += checkbox(option.name, m.choice == i)
	}
	s += "\n" + SUBTLE_STYLE.Render("use arrow keys and enter key to select.")
	return s
}

type WindowSelectorOption struct {
	name string
	on   func(model) model
}

func checkbox(label string, checked bool) string {
	if checked {
		return CHECKBOX_STYLE.Render(label) + "\n"
	} else {
		return label + "\n"
	}
}
