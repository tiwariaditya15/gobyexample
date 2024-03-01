package main 

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received:", j)
			} else {
				fmt.Println("Received all the jobs.")
				done<- true
				return
			}
		}
	}()


	for i := 0;i <= 3; i++ {
		jobs<- i
		fmt.Println("Sent job to jobs chan", i)
	}

	close(jobs)
	fmt.Println("Sent all jobs")

	<-done
	
	// _, ok:= <-jobs
	// fmt.Println("Received more jobs", ok)
}