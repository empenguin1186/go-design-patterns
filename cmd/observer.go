package main

import "github.com/empenguin1186/go-design-patterns/patterns/observer"

func main() {
	numberManager := observer.NewRandomNumberGenerator()
	generator := observer.NewNumberGenerator(numberManager)
	generator.AddObserver(observer.NewDigitObserver())
	generator.AddObserver(observer.NewGraphObserver())
	generator.Execute()
}
