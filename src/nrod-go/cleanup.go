package main

import (
	"log"
	"os"
	"os/signal"
)

func cleanUpOnInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Println("Cleaning-up...")
			cleanUp()
			log.Printf("Cleaned-up; exiting.")
			os.Exit(1)
		}
	}()
}

func cleanUp() {
	for subscriptionName, subscription := range subscriptions {
		if subscription != nil {
			log.Printf("Unsubscribing from subscription %s...", subscriptionName)
			subscription.Unsubscribe()
			log.Printf("Successfully unsubscribed from subscription %s.", subscriptionName)

			delete(subscriptions, subscriptionName)
			subscriptionWaitGroup.Done()
		}
	}

	if connection != nil {
		log.Printf("Disconnecting from connection...")
		connection.Disconnect()
		log.Printf("Sucessfully disconnected from connection.")
	}
}
