package main

import (
	"sync"
	"fmt"
)

type Container struct {
	mu sync.Mutex
	counters map[string]int
}

func (c *Container) increment(name string) {
	c.mu.Lock()
	c.counters[name]++
	defer c.mu.Unlock()
}

func main() {
	c := Container{ counters: map[string]int{"a": 0, "b": 0} }
	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {
		for i:= 1;i <= n; i++ {
			c.increment(name)
		}
		wg.Done()
	}

	wg.Add(3)
	
	go doIncrement("a", 1000)
	go doIncrement("b", 1000)
	go doIncrement("a", 1000)
	
	wg.Wait()
	fmt.Println(c.counters["a"])
	fmt.Println(c.counters["b"])
}