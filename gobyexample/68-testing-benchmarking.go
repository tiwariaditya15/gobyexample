package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func TestIntMinBasic(t *testing.T) {
	ans := IntMin(10, 4)
	if ans != 4 {
		t.Errorf("IntMin(10, 4) should be 4 got %d", ans)
	}
}

func main() {
	fmt.Println(">>>", IntMin(10, 4))
}