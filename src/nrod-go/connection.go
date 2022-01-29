package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/go-stomp/stomp/v3"
)

const STOMP_SERVER_TIMEOUT = 10 * time.Second

var connection *stomp.Conn

var stompServer string
var stompUsername string
var stompPassword string

func init() {
	stompServer = os.Getenv("STOMP_SERVER")
	stompUsername = os.Getenv("STOMP_USERNAME")
	stompPassword = os.Getenv("STOMP_PASSWORD")
}

func connect() error {
	log.Printf("Connecting: %s ...\n", stompServer)
	networkConnection, networkConnectionError := net.DialTimeout("tcp", stompServer, STOMP_SERVER_TIMEOUT)
	if networkConnectionError != nil {
		return networkConnectionError
	}

	login := stomp.ConnOpt.Login(stompUsername, stompPassword)
	newConnection, connectionError := stomp.Connect(networkConnection, login)
	connection = newConnection
	if connectionError != nil {
		return connectionError
	}

	return nil
}
