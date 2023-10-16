package linked_list

import "fmt"

func aaa(x, y int) int {
	return x + y
}

type Node struct {
	next  *Node
	value int
}
type LL struct {
	root *Node
}

func NewLL() LL {
	return LL{}
}

func (ll *LL) Print() {
	node := ll.root
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

// Push new element to the end of the queue
func (ll *LL) Push(value int) {
	if ll.root == nil {
		ll.root = &Node{value: value}
		return
	}

	node := ll.root
	for node.next != nil {
		node = node.next
	}
	node.next = &Node{value: value}

}

// Pop first element from the queue

func (ll *LL) Pop() {

	if ll.root == nil {
		return
	}

	temp := ll.root.next
	ll.root = temp

}
