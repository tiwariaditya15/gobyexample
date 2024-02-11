package main

import "fmt"


func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	
	fmt.Println("\n")

	var matrix [2][3]int

	for i := 0;i < 2; i++ {
		for j := 0;j < 3; j++ {
			matrix[i][j] = i + j
		}
	}

	fmt.Println(matrix)
}