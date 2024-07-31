package main

import (
	"os"
	"fmt"
)

func main() {
	f, err := os.CreateTemp("", "sample")
	if err != nil {
		panic(err)
	}

	defer os.Remove(f.Name())

	fmt.Println("Name:", f.Name())
}