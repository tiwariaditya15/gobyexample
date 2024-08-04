package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	os.Setenv("mode", "production")
	fmt.Println("mode>>>", os.Getenv("mode"))
	fmt.Println("mod>>>", os.Getenv("mod"))
	// fmt.Println(os.Environ())

	for i, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair)
		fmt.Println("i>>>", i)
	}
}