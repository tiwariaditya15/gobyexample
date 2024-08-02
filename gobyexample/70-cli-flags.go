package main 

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foobar", "a string")

	numbPtr := flag.Int("numb", 44, "an int")

	forkPtr := flag.Bool("fork", false, "a bool")

	var variable string

	flag.StringVar(&variable, "svar", "default", "a string")

	// to execute command line args
	flag.Parse()

	fmt.Println(*wordPtr)
	fmt.Println(*numbPtr)
	fmt.Println(*forkPtr)
	fmt.Println(variable)

}