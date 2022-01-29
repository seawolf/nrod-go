package main

import (
	"bytes"
	"encoding/json"
	"log"

	"bitbucket.org/sea_wolf/nrod-go/v2/nrod-go/messages"
	"github.com/go-stomp/stomp/v3"
)

func processMessages() error {
	message, messageError := getMessage()
	if messageError != nil {
		log.Printf("Error fetching messages: %v", messageError)
		return messageError
	}

	connection.Begin()
	processMessage(message)

	return nil
}

func getMessage() (*stomp.Message, error) {
	message := <-subscription.C

	return message, message.Err
}

func processMessage(msg *stomp.Message) {
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
			log.Println(msg)
		}
	}
}
