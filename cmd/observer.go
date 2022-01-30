package main

import "github.com/empenguin1186/go-design-patterns/patterns/observer"

func main() {
	generator := observer.RandomMessageGenerator{}
	slackNotifier := observer.SlackNotifier{}
	emailNotifier := observer.EmailNotifier{}
	generator.AddObserver(&slackNotifier)
	generator.AddObserver(&emailNotifier)
	generator.Execute()
}
