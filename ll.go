package linked_list

import "fmt"

type node[T any] struct { // hanldes any type, needs tighter constraint probably
	next  *node[T]
	value T
}

type LL[T any] struct { // The Type parameter should have same constraint as node(?)
	root *node[T]
}

func NewLL[T any]() *LL[T] { // return pointer to new LL
	return &LL[T]{}
}

func (ll *LL[T]) Print() {
	n := ll.root
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}
}

// Push new element to the end of the queue
func (ll *LL[T]) Push(value T) { // Pointer receiver
	if ll.root == nil {
		ll.root = &node[T]{value: value}
		return
	}

	n := ll.root
	for n.next != nil {
		n = n.next
	}
	n.next = &node[T]{value: value}

}

// Pop first element from the queue

func (ll *LL[T]) Pop() {

	if ll.root == nil {
		return
	}

	temp := ll.root.next
	ll.root = temp

}
