package main 

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p("Now:", now)
	p(time.UTC)
	p(time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC).Year())
	p(time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC).Month())
	p(time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC).Day())
}