package main

import (
	"fmt"
)


type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	val T
	next *element[T]
}

func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &element[T]{val: v}
		list.tail = list.head
	}else {
		list.tail.next = &element[T]{val:v}
		list.tail = list.tail.next
	}
}

func (list *List[T]) GetAllElements() []T {
	var elems []T
	for node := list.head; node!=nil; node = node.next {
	elems = 	append(elems, node.val)
	}
	return elems
}

func main() {
	fmt.Println("Linkedlist in go")
	fmt.Println("Implement push and get all elements functions")
	list := List[int]{}
	list.Push(1)
	list.Push(2)
	fmt.Println(list.GetAllElements())
}
