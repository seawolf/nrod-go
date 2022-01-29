package main

import (
	"log"

	messageTypes "bitbucket.org/sea_wolf/nrod-go/v2/nrod-go/messages"
	"github.com/go-stomp/stomp/v3"
)

func processMessages(subscription *stomp.Subscription) error {
	message, messageError := getMessage(subscription)
	if messageError != nil {
		return messageError
	}

	if message != nil {
		processMessage(subscription, message)
	}

	return nil
}

func getMessage(subscription *stomp.Subscription) (*stomp.Message, error) {
	message := <-subscription.C

	if message == nil {
		return nil, nil
	} else {
		return message, message.Err
	}
}

func processMessage(subscription *stomp.Subscription, subscriptionMessage *stomp.Message) {
	messages := messageTypes.Detect(subscriptionMessage.Body)

	for _, message := range messages {
		output := message.ToString()
		if output != "" {
			log.Printf("[%v] %v", subscription.Id(), output)
		}
	}
}
