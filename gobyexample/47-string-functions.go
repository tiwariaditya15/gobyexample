package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:", s.Contains("aditya", "ty"))
	p("Index:", s.Index("aditya", "ty"))
	p("Index:", s.Index("aditya", "ty"))
	p("Repeat:", s.Repeat("a", 100))
}