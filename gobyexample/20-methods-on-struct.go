package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	fmt.Println(r.width, r.height)
	return r.width * r.height
}


func main() {
	rectangle := rect{width: 10, height: 100}
	fmt.Println(rectangle.area())
}