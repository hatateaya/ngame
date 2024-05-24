package main

type Time struct {
	day  int
	hour int
}

func passTime(m model, hours int) model {
	for i := 0; i < hours; i++ {
		//m = m.status.trys(m)
		//m = tryEvents(m)
	}
	hours = 24*m.time.day + m.time.hour + hours
	m.time.day = hours / 24
	m.time.hour = hours % 24
	return m
}
