package main

import (
	"embed"
)

//go:embed folder/sample.txt 
var fileString string

//go:embed folder/sample.txt 
var fileByte []byte

func main() {
	print(fileString)
	print(string(fileByte))
}