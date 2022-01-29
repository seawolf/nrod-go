package main

import (
	"fmt"
	"os"

	"github.com/go-stomp/stomp/v3"
)

var subscription *stomp.Subscription

func subscribe() error {
	feedName := fmt.Sprintf("/topic/%s", subscriptionName())

	newSubscription, subscriptionError := connection.Subscribe(feedName, stomp.AckAuto)
	subscription = newSubscription

	return subscriptionError
}

func subscriptionName() string {
	return os.Getenv("FEED_TRAIN_MOVEMENTS")
}