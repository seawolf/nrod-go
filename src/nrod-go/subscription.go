package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/go-stomp/stomp/v3"
)

var subscriptionWaitGroup sync.WaitGroup
var subscriptions map[string]*stomp.Subscription
var subscriptionNames []string

func init() {
	subscriptionNames = strings.Split(os.Getenv("FEEDS_TRAIN_MOVEMENTS"), ",")
	subscriptions = map[string]*stomp.Subscription{}
}

func subscribe(subscriptionName string) (*stomp.Subscription, error) {
	feedName := fmt.Sprintf("/topic/%s", subscriptionName)
	return connection.Subscribe(feedName, stomp.AckAuto)
}
