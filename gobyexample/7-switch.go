package main

import (
	"fmt"
	"time"
)

func main() {
	switch 1+1+1 {
		case 1+1:
			fmt.Println("Yeay, weekend!")
		default: 
			fmt.Println("damn, weekday!")
	}

	t := time.Now()
	switch {
		case t.Hour() < 12: 
			fmt.Println("Afternoon")
		default:
			fmt.Println("Evening")
	}

	whoAmI := func (i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		// case string:
		// 	fmt.Println("string")	
		default:
			fmt.Printf("Do not know type of %T", t)
		}

	}

	whoAmI(1)
	whoAmI(false)
	whoAmI("aditya")
	whoAmI(1.12)
}