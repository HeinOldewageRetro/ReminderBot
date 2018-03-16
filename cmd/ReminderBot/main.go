package main

import (
	"fmt"

	"bitbucket.org/Neoin/ReminderBot/internal/eventSources"
	"bitbucket.org/Neoin/ReminderBot/internal/interfaces"
	"bitbucket.org/Neoin/ReminderBot/internal/slackChannel"
)

func main() {
	//Setup Event sources
	gc := eventSources.GoogleCalender{}

	//General filter

	//Broadcast the events

	//Per Channel Filter

	//Per slack channel send
	sender := slackChannel.Slack{}

	gc.Handler(func(m interfaces.Message) {
		err := sender.Send(m)
		if err != nil {
			fmt.Println(err)
		}
	})
	gc.Start()
}
