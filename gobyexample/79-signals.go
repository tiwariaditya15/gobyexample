package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	defer fmt.Println("Does this line print?")

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println("sig received from syscall>>>", sig)
		done<- true
	}()

	fmt.Println("Awaiting to receive on done chan.")
	<-done
	fmt.Println("We can do all cleanup here like release db connections, doing book keeping, updating cache or current state.")
	fmt.Println("EOF.")
}