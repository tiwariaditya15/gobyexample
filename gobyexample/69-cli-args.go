package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	commmandLineArgs := os.Args[1:]
	fmt.Println(commmandLineArgs)
}