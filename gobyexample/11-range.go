package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start!")
	// infos := map[string]int{"year": 2024, "date": 11}

	for i, c := range "infos" {
		fmt.Println(i, c)
	}
}