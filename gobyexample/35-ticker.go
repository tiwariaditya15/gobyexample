package main

import (
	"fmt"
	"time"
)

func main()  {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func ()  {
		for {
			select {
			case <-done:
				return
			case val := <-ticker.C:
				fmt.Println("Val:", val)
			}
		}
	}()
	
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done<- true
	fmt.Println("EOF.")
}