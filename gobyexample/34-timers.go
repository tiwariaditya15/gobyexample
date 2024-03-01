package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("timer1 fired.")

	timer2 := time.NewTimer(5 * time.Second)
	go func ()  {
		<-timer2.C
		fmt.Println("timer2 fired.")
	}()
	// Since timer2 needs 5 seconds the goroutine started and eof reached where we stopped the timer2
	timer2.Stop()
	time.Sleep(2 * time.Second)
	fmt.Println("End of file")
}