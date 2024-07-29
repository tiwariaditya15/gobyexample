package main

import (
	"bufio"
	"os"
	"fmt"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	str := []byte("hello golang\nI am js dev")
	err := os.WriteFile("/tmp/myself.txt", str, 0644)
	check(err)
	fmt.Println("Write op completed.")

	f, err1 := os.Create("/tmp/myself1.txt")
	check(err1)

	defer f.Close()

	d2 := []byte("Aditya is UI engineer")
	n2, err2 := f.Write(d2)
	check(err2)
	fmt.Printf("wrote %d bytes\n", n2)


	f.WriteString(" and knows bit of node+go\n")

	// writes to stable storage
	f.Sync()

	w := bufio.NewWriter(f)
	w.WriteString(" and this is buffered write here using bufio. :)\n")

	f2, err2 := os.Create("./docOne.txt")
	check(err2)
	w2 := bufio.NewWriter(f2)
	w2.WriteString(" and this comes from w2")

	w2.Flush()


	// to ensure all buffered operations have been applied to underlying writer
	w.Flush()

	// fmt.Println(os.Getwd())
}



/** 
		You can use following things to write to storage
			1. os
			2. bufio
*/