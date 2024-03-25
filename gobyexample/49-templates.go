package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))
	t1.Execute(os.Stdout, "blah!")
	t1.Execute(os.Stdout, 100)
	t1.Execute(os.Stdout, []int{1, 2, 3, 4})
	t1.Execute(os.Stdout, map[int]int{1:2, 2:5, 3:43, 4:2})


	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "What's good {{.}}?\n")
	t2.Execute(os.Stdout, []string{"Go", "Rust", "C#"})
}