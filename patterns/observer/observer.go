package observer

import (
	"fmt"
	"math/rand"
	"time"
)

type Observer interface {
	Update(generator *NumberGenerator)
}

type NumberGenerator struct {
	observers []Observer
	GetNumber func() int
	Execute   func()
}

func NewNumberGenerator(numberManager NumberManager) *NumberGenerator {
	return &NumberGenerator{observers: []Observer{}, GetNumber: numberManager.GetNumber, Execute: numberManager.Execute}
}

func (g *NumberGenerator) AddObserver(observer Observer) {
	g.observers = append(g.observers, observer)
}

func (g *NumberGenerator) NotifyObservers() {
	for _, e := range g.observers {
		e.Update(g)
	}
}

type NumberManager interface {
	GetNumber() int
	Execute()
}

type RandomNumberGenerator struct {
	number          int
	notifyObservers func()
}

func NewRandomNumberGenerator() *RandomNumberGenerator {
	return &RandomNumberGenerator{number: 0, notifyObservers: NumberGenerator{}.NotifyObservers}
}

func (r *RandomNumberGenerator) GetNumber() int {
	return r.number
}

func (r *RandomNumberGenerator) Execute() {
	for i := 0; i < 20; i++ {
		r.number = rand.Intn(100)
		r.notifyObservers()
	}
}

type DigitObserver struct{}

func NewDigitObserver() *DigitObserver {
	return &DigitObserver{}
}

func (d *DigitObserver) Update(generator *NumberGenerator) {
	fmt.Printf("DigitObserver:%d", generator.GetNumber())
	time.Sleep(time.Millisecond * 100)
}

type GraphObserver struct{}

func NewGraphObserver() *GraphObserver {
	return &GraphObserver{}
}

func (g *GraphObserver) Update(generator *NumberGenerator) {
	fmt.Println("GraphObserver")
	count := generator.GetNumber()
	for i := 0; i < count; i++ {
		fmt.Print("*")
	}
	fmt.Println("*")
	time.Sleep(time.Millisecond * 100)
}
