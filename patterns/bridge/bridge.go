package bridge

import (
	"fmt"
)

type DisplayImpl interface {
	rawOpen()
	rawPrint()
	rawClose()
}

type Display struct {
	DisplayImpl
}

func (d *Display) Open() {
	d.rawOpen()
}

func (d *Display) Print() {
	d.rawPrint()
}

func (d *Display) Close() {
	d.rawClose()
}

func (d *Display) Display() {
	d.Open()
	d.Print()
	d.Close()
}

func NewDisplay(impl DisplayImpl) *Display {
	return &Display{impl}
}

type CountDisplay struct {
	*Display
}

func NewCountDisplay(impl DisplayImpl) *CountDisplay {
	return &CountDisplay{NewDisplay(impl)}
}

func (d *CountDisplay) MultiDisplay(times int) {
	d.Open()
	for i := 0; i < times; i++ {
		d.Print()
	}
	d.Close()
}

type StringDisplayImpl struct {
	text  string
	width int
}

func NewStringDisplayImpl(text string) *StringDisplayImpl {
	return &StringDisplayImpl{text: text, width: len(text)}
}

func (s *StringDisplayImpl) rawOpen() {
	s.printLine()
}

func (s *StringDisplayImpl) rawPrint() {
	fmt.Println("|" + s.text + "|")
}

func (s *StringDisplayImpl) rawClose() {
	s.printLine()
}

func (s *StringDisplayImpl) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
