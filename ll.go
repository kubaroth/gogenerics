package linked_list

import "fmt"

type Node[T any] struct { // hanldes any type, needs tighter constraint probably
	next  *Node[T]
	value T
}

type LL[T any] struct { // The Type parameter should have same constraint as Node(?)
	root *Node[T]
}

func NewLL[T any]() *LL[T] { // return pointer to new LL
	return &LL[T]{}
}

func (ll *LL[T]) Print() {
	node := ll.root
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

// Push new element to the end of the queue
func (ll *LL[T]) Push(value T) { // Pointer receiver
	if ll.root == nil {
		ll.root = &Node[T]{value: value}
		return
	}

	node := ll.root
	for node.next != nil {
		node = node.next
	}
	node.next = &Node[T]{value: value}

}

// Pop first element from the queue

func (ll *LL[T]) Pop() {

	if ll.root == nil {
		return
	}

	temp := ll.root.next
	ll.root = temp

}
