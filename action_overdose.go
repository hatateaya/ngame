package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ActionOverdose struct {
}

func (this ActionOverdose) available(m model) bool {
	return true
}

func (this ActionOverdose) do(m model) model {
	m = passTime(m, 1)
	m.window = new(ActionOverdose)
	return m
}

func (this ActionOverdose) update(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	return this.selector(m).update(m, msg)
}

func (this ActionOverdose) view(m model) string {
	return "服用什么药物呢？\n\n" + this.selector(m).view(m)
}

func (this ActionOverdose) selector(m model) WindowSelector {
	selector := new(WindowSelector)
	if m.data.med_dxm == 0 {
		selector.options = append(selector.options,
			WindowSelectorOption{"氢溴酸右美沙芬片",
				func(m model) model {
					dialog := new(WindowShowMsg)
					dialog.msg = "您服用了十二片氢溴酸右美沙芬片"
					m.dialog = dialog
					return m
				}})
	}
	if m.data.med_bron != 0 {

	}
	if m.data.med_bzd != 0 {

	}
	if m.data.med_bzd != 0 {

	}
	if m.data.med_dph != 0 {

	}
	if m.data.med_admt != 0 {

	}
	if m.data.med_pr80 != 0 {

	}
	selector.options = append(selector.options,
		WindowSelectorOption{"返回",
			func(m model) model { m.window = new(WindowAction); return m }})
	return *selector
}
