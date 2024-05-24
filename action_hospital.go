package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ActionHospital struct {
}

func (this ActionHospital) available(m model) bool {
	return true
}

func (this ActionHospital) do(m model) model {
	m = passTime(m, 1)
	m.window = new(ActionHospital)
	return m
}

func (this ActionHospital) update(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	return this.selector(m).update(m, msg)
}

func (this ActionHospital) view(m model) string {
	return "服用什么药物呢？\n\n" + this.selector(m).view(m)
}

func (this ActionHospital) selector(m model) WindowSelector {
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
	selector.options = append(selector.options,
		WindowSelectorOption{"返回",
			func(m model) model { m.window = new(WindowAction); return m }})
	return *selector
}
