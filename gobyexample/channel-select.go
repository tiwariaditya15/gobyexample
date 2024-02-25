package main

import (
	"fmt"
	"time"
)

func fetchUser(c1 chan<- string) { 
	time.Sleep(2 * time.Second)
	c1 <- "Completed fetching user."
}

func fetchRepos(c2 chan<- string) {
	time.Sleep(1 * time.Second)
	c2 <- "Completed fetching repos based on User."
 }

func fetchRepoData(c3 chan<- string) {
	time.Sleep(1 * time.Second)
	c3 <- "Completed fetching repo data after fetch repo based on user."
}

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c3 := make(chan string, 1)

	go fetchUser(c1)
	go fetchRepos(c2)
	go fetchRepoData(c3)

	for i := 0;i < 3; i++ {
		select {
		case msg := <-c1: 
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)	
		case msg := <-c3:
			fmt.Println(msg)
		}
	}
}