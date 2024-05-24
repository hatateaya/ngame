package main

type Event interface {
	try(model) model
}

func tryEvents(m model) model {
	for _, event := range m.events {
		m = event.try(m)
	}
	return m
}
