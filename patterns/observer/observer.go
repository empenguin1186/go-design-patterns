package observer

import (
	"fmt"
	"math/rand"
)

type Observer interface {
	Notify(generator MessageGenerator)
}

type MessageGenerator interface {
	AddObserver(observer Observer)
	DeleteObserver(observer Observer)
	GetMessage() string
	Execute()
}

type RandomMessageGenerator struct {
	observers []Observer
	message   string
}

func (r *RandomMessageGenerator) AddObserver(observer Observer) {
	r.observers = append(r.observers, observer)
}

func (r *RandomMessageGenerator) DeleteObserver(observer Observer) {
	var observers []Observer
	for _, e := range r.observers {
		if e != observer {
			observers = append(observers, e)
		}
	}
	r.observers = observers
}

func (r *RandomMessageGenerator) GetMessage() string {
	return r.message
}

func (r *RandomMessageGenerator) Execute() {
	r.message = fmt.Sprintf("message changed. number=%d", rand.Intn(100))
	for _, e := range r.observers {
		e.Notify(r)
	}
}

type SlackNotifier struct{}

func (s *SlackNotifier) Notify(generator MessageGenerator) {
	fmt.Printf("send message to slack. message = \"%s\"\n", generator.GetMessage())
}

type EmailNotifier struct{}

func (e *EmailNotifier) Notify(generator MessageGenerator) {
	fmt.Printf("send message to email. message = \"%s\"\n", generator.GetMessage())
}
