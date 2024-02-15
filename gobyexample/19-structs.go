package main

import "fmt"

type sport struct {
	name string
	country string
	level int
}

func main() {
	aditya := sport{country: "Australia", name: "aditya", level: 4}
	fmt.Println(aditya.name)	
	
	newAditya := &aditya
	newAditya.name =  "Aditya Tiwari"
	fmt.Println(newAditya.name)	
	fmt.Println(aditya.name)	

	inlineStructure := struct {
		isOnline bool
		devs int
	} {
		false,
		6,
	}

	fmt.Println(inlineStructure)
}