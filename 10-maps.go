package main

import (
	"fmt"
	"maps"
)


func main() {
	marks := make(map[string]int)

	marks["english"] = 84
	marks["maths"] = 76
	english := "englishh"
	englishmatrks, present := marks[english]
	fmt.Println(englishmatrks, present)
	if !present {
		fmt.Println("englishh key not in map")
	}
	fmt.Println("len of map marks:", len(marks))
	
	clear(marks)
	fmt.Println(marks[english])
	fmt.Println("len of map marks:", len(marks))
	fmt.Println(marks)

	m1 := map[string]int{"name": 11, "age": 22}
	m2 := map[string]int{"nameK": 11, "age": 22}

	if maps.Equal(m1, m2) {
		fmt.Println("m1 and m2 hold same values")
	}
	delete(m2, "nameK")
	fmt.Println(m2)
}

