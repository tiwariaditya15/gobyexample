package main

import "fmt"

func passByValue(value int) {
	value = 34
}

func passByRef(value *int) {
	fmt.Println(value, *value)
	*value = 34
}

func main() {
	name := "aditya"
	age := 24
	fmt.Println("Age initially - ", age)
	passByValue(age)
	fmt.Println("Age after passByValueCall - ", age)
	passByRef(&age)
	fmt.Println("Age after passByRef - ", age)
	fmt.Println("&age:", &age)
	fmt.Println("&name:", &name)

	for _, char := range "aditya" {
		fmt.Println(char)
	}

}