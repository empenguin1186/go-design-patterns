package main

import (
	"github.com/empenguin1186/go-design-patterns/patterns/observer"
	"time"
)

func main() {
	generator := observer.RandomMessageGenerator{Ch: make(chan string)}
	slackNotifier := observer.SlackNotifier{}
	emailNotifier := observer.EmailNotifier{}
	generator.AddObserver(&slackNotifier)
	generator.AddObserver(&emailNotifier)
	generator.Execute()

	go func() {
		time.Sleep(time.Second * 2)
		generator.SetMessage("hoge")
	}()

	go func() {
		time.Sleep(time.Second * 3)
		generator.SetMessage("fuga")
	}()
}
