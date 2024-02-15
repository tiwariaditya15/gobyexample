package main

import (
	"fmt"
)

type List[T any] struct {
	head, tail *item[T]
}

type item[T any] struct {
	val T
	next *item[T]
}

func MapKeys[K comparable, V any] (m map[K]V) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}

	return r
}

func (lst *List[T]) Push(val T) {
	if lst.head == nil {
		lst.head = &item[T]{val: val}
		lst.tail = lst.head
	} else {
		lst.tail.next = &item[T]{val: val}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var r []T

	iter := lst.head
	for ;iter != nil; {
		r = append(r, iter.val)
		iter = iter.next
	}

	return r
}

func main() {
	z := map[string]string{"name": "aditya", "age": "24"}
	fmt.Println(MapKeys(z))

	lst := List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)
	fmt.Println(lst.GetAll())
}