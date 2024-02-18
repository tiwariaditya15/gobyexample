package main 

import "fmt"

func main() {
	messages := make(chan string)

	go func () { messages <- "ping" }()

	ret :=  "<-messages"
	fmt.Println(ret)
}