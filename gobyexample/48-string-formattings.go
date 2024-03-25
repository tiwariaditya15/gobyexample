package main

import (
	"fmt"
)

type point struct {
	fname string
	lname string
}

func main() {
	info := point{"aditya", "tiwari"}
	fmt.Printf("point: %v\n", info)
	fmt.Printf("point: %+v\n", info)
	fmt.Printf("point: %#v\n", info)
	fmt.Printf("point type: %T\n", info)
}