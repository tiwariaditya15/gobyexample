package main

import (
	"os"
	"fmt"
	"io/fs"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Makes directory
	os.Mkdir("subdir", 0755)
	// Removes previous directory
	// defer os.RemoveAll("subdir")


	createEmptyFile := func (name string) {
		d := []byte("Go go go!")
		check(os.WriteFile(name, d, 0755))
	}

	createEmptyFile("subdir/temp.txt")


	err := os.MkdirAll("subdir/parent/child", 0755)
	check(err)


    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

	fmt.Println("Listing dir > ")
	c, err := os.ReadDir("subdir")
	check(err)
	for _, entry := range c {
		entry.Name()
		// fmt.Println("Name:", entry.Name(), "isDirectory:", entry.IsDir())
	}

	err1 := os.Chdir("subdir/parent/child")
	check(err1)

	err2 := os.Chdir("../../..")
	check(err2)
	
	err3 := filepath.WalkDir("subdir", visit)
	check(err3)
}

func visit(path string, d fs.DirEntry, e error) error {
	if e != nil {
		return e
	}

	fmt.Println("Path:", path, "Name:", d.Name(), "isDir:", d.IsDir())
	return nil
}