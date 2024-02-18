package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done...")
	
	done <- false
}

func main() {
	done := make(chan bool, 1)

	go worker(done)

	workerResult := <-done
	if workerResult == true {
		fmt.Println("Worker completed task.")	
	} else {
		fmt.Println("Worker failed task.")
	}
}