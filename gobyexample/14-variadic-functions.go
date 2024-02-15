package main

import "fmt"

func sum(marks ...int) {
	fmt.Println(marks)
	for index, mark := range marks {
		fmt.Println(mark, " at ", index)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	// arr = make([]int, 8)
	// append(arr, 6)
	sum(arr...)
}