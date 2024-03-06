package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops atomic.Uint64
	var count int
	var wg sync.WaitGroup 

	for i := 1;i <= 50;i++ {
		wg.Add(1)
		go func () {
			for j := 1;j <= 1000;j++ {
				ops.Add(1)
				count = count + 1
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("OPS:", ops.Load())
	fmt.Println("count:", count)
	fmt.Println("EOF.")
}