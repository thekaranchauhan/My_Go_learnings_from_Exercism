package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	hour := (h + m/60) % 24
	minute := m % 60
	for hour < 0 {
		hour += 24
	}
	for minute < 0 {
		minute += 60
		hour -= 1
	}
	return Clock{hour, minute}
}

func (c Clock) Add(m int) Clock {
	c.minute += m
	for c.minute >= 60 {
		c.hour += 1
		c.minute -= 60
	}
	for c.hour >= 24 {
		c.hour -= 24
	}
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.minute -= m
	for c.minute < 0 {
		c.hour--
		c.minute += 60
	}
	for c.hour < 0 {
		c.hour += 24
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
