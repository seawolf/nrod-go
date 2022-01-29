package main

import (
	"bytes"
	"encoding/json"
	"log"

	"bitbucket.org/sea_wolf/nrod-go/v2/nrod-go/messages"
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

func processMessage(subscription *stomp.Subscription, msg *stomp.Message) {
	var messages []messages.MovementMessage
	err := json.Unmarshal(msg.Body, &messages)

	if err != nil {
		var out bytes.Buffer
		json.Indent(&out, msg.Body, "", "\t")
		log.Printf("error parsing JSON: %v\n%s", err, out.String())
	}

	for _, message := range messages {
		msg := message.ToString()
		if msg != "" {
			log.Printf("[%v] %v", subscription.Id(), msg)
		}
	}
}
