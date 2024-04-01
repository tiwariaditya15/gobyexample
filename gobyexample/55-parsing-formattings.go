package main


import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p(now)
	p(now.Format(time.RFC3339))
}