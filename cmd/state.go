package main

import (
	"fmt"
)

type DayOfWeek int

const (
	Sunday DayOfWeek = iota
	Monday
	Tuesday
	//Sunday = DayOfWeek("Sunday")
	//Monday = DayOfWeek("Monday")
	//Tuesday = DayOfWeek("Monday")
)

func main() {
	//context := state.NewContext()
	hour := 18
	dayOfWeek := Sunday

	for i := 0; i < 4; i++ {
		printStoreIsOpen(dayOfWeek, hour)
		if dayOfWeek == 2 {
			dayOfWeek = 0
		} else {
			dayOfWeek++
		}
	}

}

func printStoreIsOpen(dayOfWeek DayOfWeek, hour int) {
	var dayOfWeekString string
	var openOrClose string
	switch dayOfWeek {
	case Sunday:
		dayOfWeekString = "Sunday"
		if hour >= 10 && hour <= 17 {
			openOrClose = "open"
		} else {
			openOrClose = "close"
		}
		break
	case Monday:
		dayOfWeekString = "Monday"
		openOrClose = "close"
		break
	case Tuesday:
		dayOfWeekString = "Tuesday"
		if hour >= 9 && hour <= 21 {
			openOrClose = "open"
		} else {
			openOrClose = "close"
		}
		break
	default:
		dayOfWeekString = "Undefined"
		openOrClose = "Undefined"
	}
	fmt.Printf("Today is %s. Store %s.\n", dayOfWeekString, openOrClose)
}
