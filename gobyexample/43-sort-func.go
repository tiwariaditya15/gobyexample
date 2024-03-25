package main 

import (
	"slices"
	"fmt"
	"cmp"
)

func main() {
	fruits := []string{"watermelon", "banana", "grapes"}

	compare := func (a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	
	slices.SortFunc(fruits, compare)

	fmt.Println(fruits)

	// Custom sort on type that aren't built-in
	type Person struct {
		name string
		age int
	}

	peoples := []Person{
		Person{name: "pa", age: 33},
		Person{name: "ad", age: 24},
		Person{name: "pr", age: 31},
	}

	slices.SortFunc(peoples, func (a, b Person) int {
		return cmp.Compare(b.age, a.age)
	})

	fmt.Println(peoples)
}