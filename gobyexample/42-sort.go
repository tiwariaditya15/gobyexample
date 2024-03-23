package main


import (
	"fmt"
	"slices"
)

func main() {
	s := []int{6, 5, 21, 1, 23, 3}

	fmt.Println(s)
	fmt.Println(slices.IsSorted(s))
	slices.Sort(s)
	fmt.Println(s)
 }