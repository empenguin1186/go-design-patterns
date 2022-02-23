package main

import "github.com/empenguin1186/go-design-patterns/patterns/bridge"

func main() {
	d1 := bridge.NewDisplay(bridge.NewStringDisplayImpl("Hello, Japan."))
	d2 := bridge.NewCountDisplay(bridge.NewStringDisplayImpl("Hello, World."))

	d1.Display()
	d2.MultiDisplay(5)
}
