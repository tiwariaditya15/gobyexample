package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "index.html")
	fmt.Println(p)

	fmt.Println(filepath.Join("dir1//", "index.tsx"))
	fmt.Println(filepath.Join("dir1/../", "index.tsx"))

	fmt.Println(filepath.Dir(p))
	fmt.Println(filepath.Base(p))

	fmt.Println(filepath.IsAbs("/dir1/dir2/index.tsx"))
	fmt.Println(filepath.IsAbs("dir1/dir2/index.tsx"))
	filename := "config.tsx"
	fileExt := filepath.Ext(filename)
	fmt.Println(fileExt)
	fmt.Println(strings.TrimSuffix(filename, fileExt))
}
