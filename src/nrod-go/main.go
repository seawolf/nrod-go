package main

import (
	"log"
)

func main() {
	cleanUpOnInterrupt()

	connectionError := connect()
	defer connection.Disconnect()
	if connectionError != nil {
		cleanUp()
		log.Fatalf("error establishing connection: %v", connectionError.Error())
	} else {
		log.Printf("Connected: %v", connection.Session())
	}

	subscriptionError := subscribe()
	if subscriptionError != nil {
		cleanUp()
		log.Printf("error subscribing: %v", subscriptionError)
	} else {
		log.Printf("Subscribed: %v", subscription.Id())
	}

	var processError error

	for {
		log.Println("Waiting for messages...")

		processError = processMessages()

		if processError != nil {
			cleanUp()
			log.Printf("error processing: %v", processError)
		}
	}
}
