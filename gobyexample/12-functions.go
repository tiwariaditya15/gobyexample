package main

import "fmt"

func adder(a, b int) int {  // important to put return type
	return a + b;
}

func main() {
	sum := adder(1, 2)
	fmt.Println(sum)
}