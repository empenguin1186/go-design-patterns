package state

import "fmt"

var sundaySingleton *sunday = &sunday{}
var mondaySingleton *monday = &monday{}
var tuesdaySingleton *tuesDay = &tuesDay{}

type context struct {
	dayOfWeekStates []DayOfWeekState
	index           int
}

func NewContext() *context {
	return &context{dayOfWeekStates: []DayOfWeekState{sundaySingleton, mondaySingleton, tuesdaySingleton}, index: 0}
}

func (c *context) OneDayPass() {
	if c.index == len(c.dayOfWeekStates)-1 {
		c.index = 0
	} else {
		c.index++
	}
}

func (c *context) GetMessage(hour int) string {
	dayOfWeek := c.dayOfWeekStates[c.index]
	var text string
	if dayOfWeek.isStoreOpen(hour) {
		text = "is open"
	} else {
		text = "is close"
	}
	return fmt.Sprintf("Today is %s. Store %s.", dayOfWeek.ToString(), text)
}

type DayOfWeekState interface {
	ToString() string
	isStoreOpen(hour int) bool
}

type sunday struct{}

func (s *sunday) ToString() string {
	return "Sunday"
}

func (s *sunday) isStoreOpen(hour int) bool {
	if hour >= 10 && hour <= 17 {
		return true
	}
	return false
}

type monday struct{}

func (m *monday) ToString() string {
	return "Monday"
}

func (m *monday) isStoreOpen(hour int) bool {
	return false
}

type tuesDay struct{}

func (t *tuesDay) ToString() string {
	return "TuesDay"
}

func (t *tuesDay) isStoreOpen(hour int) bool {
	if hour >= 9 && hour <= 21 {
		return true
	}
	return false
}
