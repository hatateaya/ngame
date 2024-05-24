package main

import tea "github.com/charmbracelet/bubbletea"

type WindowView struct {
}

func (this WindowView) update(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	return this.selector().update(m, msg)
}

func (this WindowView) view(m model) string {
	return "请查看数据：\n\n" + this.selector().view(m)
}

func (this WindowView) selector() WindowSelector {
	selector := new(WindowSelector)
	selector.options = []WindowSelectorOption{
		{"返回选项", func(m model) model { m.window = new(WindowAction); return m }}}
	return *selector
}
