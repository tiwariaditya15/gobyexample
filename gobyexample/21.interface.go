package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}


func (r rect) area() float64 {
	return r.width*r.height;
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height;
}


func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius;
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius;
}

func printGeometries(g geometry) {
	fmt.Println("area - ", g.area(), "perim - ", g.perim())
}


func main() {
	rectangle := rect{width: 10, height: 40}
	round := circle{radius: 10}
	printGeometries(rectangle)
	printGeometries(round)
}