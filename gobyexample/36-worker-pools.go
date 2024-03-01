package main

import (
	"fmt"
	"time"
)

/**
The workers are spawned as goroutines so they run parallely.
So whole task ie 5 and execution time ie 2s per task would've taken (5*2)s but goroutines make it run parallel and complete faster.
*/

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, "started job ", j)
		time.Sleep(2 * time.Second)
		fmt.Println("worker ", id, "finished job ", j)
		results<- j*2
	}
	fmt.Println("Closing worker:", id)
}

func main()  {
	numOfJobs := 5
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	// Start workers
	for i := 1;i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 1;j <= numOfJobs; j++ {
		jobs<- j
	}

	for k := 1;k <= numOfJobs; k++ {
		fmt.Println(<-results)
	}

	close(jobs)
	time.Sleep(2 * time.Second)
	fmt.Println("EOF.")
}