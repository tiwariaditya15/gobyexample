package main 

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)


func main() {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			fmt.Println(strings.ToUpper(scanner.Text()), "eof goes where?")
		}


		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
}