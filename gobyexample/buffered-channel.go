package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	
	messages <- "buffering messages"
	messages <- "here even if channel isnt ready to receive"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// this is failing and throwinfg fatal error deadlock since loop will keep on waiting to receive more msg but there are no active goroutines, hence dead lock
	// for msg := range messages {
	// 	fmt.Println(msg)
	// }
}