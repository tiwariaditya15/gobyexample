package main

import "fmt"

func multiReturns(a, b int) (int, int) {
	return a+1, b+2
}

func main() {
	a, b := multiReturns(1, 2)
	fmt.Println(a)
	fmt.Println(b)
}
