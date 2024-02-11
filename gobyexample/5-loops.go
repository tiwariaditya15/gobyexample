package main

import (
	"fmt"
)

func main() {
	// Sort of while loop
	fmt.Println("While loop")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("\n")
	
	fmt.Println("For loop")
	for i := 1;i <= 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("\n")

	j := 1
	for {

		j = j + 1

		if j == 2 {
			continue
		}

		fmt.Println("Woot!", j)

		if j > 3 {
			break
		}
	}
}