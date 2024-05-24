package main

type Action interface {
	available(m model) bool
	do(m model) bool
}
