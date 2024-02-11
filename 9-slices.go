package main

import (
	"slices"
	"fmt"
)

func main() {
	var ages []int // a slice which doesn't need to specify the number of elements
	ages = make([]int, 6)
	fmt.Println("len of ages:", len(ages))
	fmt.Println("\nages:", ages)
	
	ages = append(ages, 30, 1)
	
	fmt.Println("\nages:", ages)

	fortunes := make([]int, len(ages))
	copy(fortunes, ages)

	fmt.Println(fortunes)

	l := ages[:6]
	fmt.Println("l:", l)

	t2 := []string{"a", "b", "d"}
	t3 := []string{"a", "b", "", "d"}


	if slices.Equal(t2, t3) {
		fmt.Println("t2 and t3 are equal")
	} else {
		fmt.Println("t2 and t3 aren't equal")
			
	}


	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}

/**
	make()
	append()
	copy()
	slices.Equal

**/

