package main

import (
	"fmt"
	"math/rand/v2"
)


func main() {
	p := fmt.Println
	p(rand.IntN(2))
	p(rand.Float64())
	p((rand.Float64() * 5) + 5, ",")
	p((rand.Float64() * 5) + 5)


	// You can have a seed which guarantess on passing 2 same numbers, you will get exactly same output
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	p()
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
}