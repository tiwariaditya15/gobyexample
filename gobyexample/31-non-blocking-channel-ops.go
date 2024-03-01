package main

import (
	"fmt"
)

func main() {
	// non buffered channels
	messages := make(chan string)
	signals := make(chan bool)

	// Non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println(msg)
	default:
		fmt.Println("No messages received!")
	}

	msg := "hello there!"
	// Non-blocking send
	select {
	case messages<- msg:
		fmt.Println("Sent to messages!")
	default:
		fmt.Println("Message channel not ready to accept.")
	}

	// multi-way-non-blocking
	select {
	case msg := <-messages:
		fmt.Println("Received on messages channel:", msg)
	case sig := <-signals:
		fmt.Println("Received on signals channel:", sig)
	default:
		fmt.Println("Default way.")
	}
}