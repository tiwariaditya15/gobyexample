package main

import (
	"fmt"
)

var a int = 1

func main() {
	fmt.Println(a)
	var a string
	a = "aditya"
	fmt.Println(a)

	var c, b int = 30, 06
	fmt.Println(c, b)

	var d string
	fmt.Println(d)

	// only inside functions
	e := "tiwari"
}