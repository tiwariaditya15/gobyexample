package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// data, err := os.ReadFile("./doc.txt")
	// check(err)
	// fmt.Println(string(data))

	f, err := os.Open("./doc.txt")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	
	
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%d bytes: %s\n", n2, string(b2[:n2]))


	o3, err := f.Seek(18, 0)
	check(err)
	b3 := make([]byte, 3)
	n3, err := io.ReadAtLeast(f, b3, 3)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3[:n3]))

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(8)
	check(err)
	fmt.Printf("%s", string(b4))
}