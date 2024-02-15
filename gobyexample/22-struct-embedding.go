package main

import "fmt"

type describer interface {
	describe() string
}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("num:", b.num)
}

type car struct {
	base
	name string
}

func main() {
	b := base{14}
	b.describe()


	c := car{
		base: base{
			num: 7,
		},
		name: "Honda",
	}

	fmt.Println(c.describe())

	var d describer
	d = c

	fmt.Println(d.describe())
}