package main

import tea "github.com/charmbracelet/bubbletea"

type WindowAction struct {
}

func (this WindowAction) update(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	return this.selector().update(m, msg)
}

func (this WindowAction) view(m model) string {
	return "请选择您要做的事情：\n\n" + this.selector().view(m)
}

func (this WindowAction) selector() WindowSelector {
	selector := new(WindowSelector)
	selector.options = []WindowSelectorOption{
		{"磕药", func(m model) model { m = new(ActionOverdose).do(m); return m }},
		{"查看数据", func(m model) model { m.window = new(WindowView); return m }}}
	return *selector
}
