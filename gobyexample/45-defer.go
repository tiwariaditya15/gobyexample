package main

// mostly used to cleaup

import (
	"os"
	"fmt"
)

func createFile(path string) *os.File {
	fmt.Println("Creating file.")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing to file.")
	fmt.Fprintln(f, "This file is defer.txt")
}
func closeFile(f *os.File) {
	fmt.Println("Closing file.")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	f := createFile("/tmp/defer.text")
	defer closeFile(f)
	writeFile(f)
	fmt.Println("EOF.")
}