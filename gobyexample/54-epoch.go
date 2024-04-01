package main 

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	
	p(now.Unix())
	p(now.UnixMilli())
	p(now.UnixNano())
}