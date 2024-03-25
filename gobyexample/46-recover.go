package main

import "fmt"

func someIOCall() {
	panic("Boom boom boom! Trata trata traaaaa!")
}

func main() {

	defer func () {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error: \n", r)
		}
	}()
	someIOCall()
	fmt.Println("EOF.")
}