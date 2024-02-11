package main

import (
	"fmt"
)

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even number")
		} else {
		fmt.Println("7 is odd number")	
	}

	if 6%2 == 0 {
		fmt.Println("6 is even number")
		} else {
		fmt.Println("6 is odd number")	
	}

	if num := 19; num < 0 {
		fmt.Println("num is negative number")
	} else if num < 10 {	
		fmt.Println("num is single digit number")
	} else {
		fmt.Println("num is double digit number")
	}
}