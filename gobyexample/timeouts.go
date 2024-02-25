package main 

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1<- "Fetch data" 
	}()

	select {
	case msg := <-c1:
	fmt.Println(msg)
	case <-time.After(1 * time.Second):
	fmt.Println("Timeout on c1")
	}


	// c2 := make(chan string, 1)
	// go func() {
	// 	time.Sleep(time.Second)
	// 	c2<- "fetched data inside f2"
	// }()

	// select {
	// case msg := <-c2:
	// fmt.Println(msg)
	// case <- time.After(2 * time.Second):
	// fmt.Println("Timeout on c2")	
	// }
}