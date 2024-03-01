package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("timer1 fired.")

	timer2 := time.NewTimer(time.Second)
	go func ()  {
		<-timer2.C
		fmt.Println("timer2 fired.")
	}()

	time.Sleep(2 * time.Second)
}