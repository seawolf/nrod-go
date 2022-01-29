package main

import (
	"log"
)

func main() {
	cleanUpOnInterrupt()

	connectionError := connect()
	defer connection.Disconnect()
	if connectionError != nil {
		log.Fatalf("error establishing connection: %v", connectionError.Error())
	} else {
		log.Printf("Connected: %v", connection.Session())

		for _, subscriptionName := range subscriptionNames {
			subscriptionWaitGroup.Add(1)
			go workSubscription(subscriptionName)
		}
		subscriptionWaitGroup.Wait()
	}

	if len(subscriptions) > 0 {
		log.Println("Cleaning up...")
		cleanUp()
	}
}

func workSubscription(subscriptionName string) {
	subscription, subscriptionError := subscribe(subscriptionName)

	if subscriptionError != nil {
		log.Printf("error subscribing: %v", subscriptionError)
		return
	}

	subscriptions[subscriptionName] = subscription
	log.Printf("Subscribed: %s (%v)", subscriptionName, subscription.Id())

	for {
		if subscription == nil {
			return
		}

		processMessages(subscription)
	}
}
