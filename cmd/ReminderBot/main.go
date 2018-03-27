package main

import (
	"fmt"

	"bitbucket.org/Neoin/ReminderBot/internal/eventSources"
	"bitbucket.org/Neoin/ReminderBot/internal/interfaces"
	"bitbucket.org/Neoin/ReminderBot/internal/slackChannel"
	"github.com/namsral/flag"
)

var meetupKey = flag.String("meetupKey", "", "Meetup API key")
var doSend = flag.Bool("s", false, "Are we sending to slack")
var slackWebHook = flag.String("slackWebHook", "", "slackWebHook")

func main() {
	flag.Parse()
	//Setup Event sources
	//gc := eventSources.GoogleCalender{}
	meetup := eventSources.Meetup{Key: *meetupKey}
	//General filter

	//Broadcast the events

	//Per Channel Filter

	//Per slack channel send
	sender := &slackChannel.Slack{SlackWebHook: *slackWebHook}

	/*
		gc.Handler(SendWrapper(sender))
		go gc.Start()
	*/

	meetup.Handler(SendWrapper(sender))
	meetup.Start()
}

func SendWrapper(sender interfaces.ChannelSender) func(m interfaces.Message) {
	return func(m interfaces.Message) {
		fmt.Println(m.String())
		if *doSend {
			err := sender.Send(m)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
