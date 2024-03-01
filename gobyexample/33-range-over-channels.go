package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)
	queue<- "First"
	queue<- "Second"
	close(queue)

	// this goes to show you can still receive values even if channel is closed
	for r := range queue {
		fmt.Println(r)
	}
}