package main

import (
	"fmt"
)

func ping(pings chan<- string, msg string) {
	pings <- msg
}
/**
1. pings <-chan: You can only receive message from this channel
2. pongs chan<-: You can only send message to this channel

This helps up right type-safe programs where you accidentally don't send/receive on channel where you're not supposed to

*/

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
	fmt.Println(<-pongs)
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "What's good gang?")
	pong(pings, pongs)

	// fmt.Println(<-pongs)
}