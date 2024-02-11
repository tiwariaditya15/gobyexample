package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start!")

	for i, c := range "infos" {
		fmt.Println(i, c)
	}
}