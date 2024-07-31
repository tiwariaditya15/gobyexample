package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func main() {
	f, err := os.CreateTemp("", "sample")
	if err != nil {
		panic(err)
	}

	defer os.Remove(f.Name())

	fmt.Println("Name:", f.Name())

	dname, err1 := os.MkdirTemp("", "tempDir")
	defer os.RemoveAll("tempDir")
	if err1 != nil {
		panic(err1)
	}
	
	fname := filepath.Join(dname, "tempDoc.txt")
	err2 := os.WriteFile(fname, []byte{1, 2}, 0666)
	if err2 != nil {
		panic(err2)
	}
}