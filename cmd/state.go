package main

import (
	"fmt"
	"github.com/empenguin1186/go-design-patterns/patterns/state"
)

func main() {
	context := state.NewContext()
	hour := 18
	for i := 0; i < 4; i++ {
		fmt.Println(context.GetMessage(hour))
		context.OneDayPass()
	}
}
