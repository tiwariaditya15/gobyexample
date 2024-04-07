package main


import (
	"fmt"
	"strconv"
)


func main() {
	_, err := strconv.ParseFloat("adi", 64)
	if err != nil {
		fmt.Println("Parsing failed.")
	}

	b10, _ := strconv.Atoi("145")
	fmt.Println(b10)

	d, _ := strconv.ParseInt("0x1c8", 0 , 64)
	fmt.Println(d)
}