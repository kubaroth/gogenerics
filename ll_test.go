package linked_list

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	value := aaa(1, 2)
	fmt.Println("TestOne")
	if value != 3 {
		t.Fatalf("ASASA")
	}
}

func TestLL(t *testing.T) {
	ll := NewLL()
	if ll.root != nil {
		t.Fatalf("LL root node is not nil")
	}
}
func TestLL_Push(t *testing.T) {
	ll := NewLL()
	ll.Push(1)
	if ll.root.value != 1 {
		t.Fatalf("LL root node is not 1")
	}

	ll.Push(2)
	ll.Push(3)
	ll.Print()
	if ll.root.next.next.value != 3 {
		t.Fatalf("LL third node is not 3")
	}
}

func TestLL_Pop(t *testing.T) {
	ll := NewLL()
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)

	ll.Pop()
	// ll.Print()

	if ll.root.value != 2 {
		t.Fatalf("LL root node is not 2")
	}
	ll.Pop()
	if ll.root.value != 3 {
		t.Fatalf("LL root node is not 2")
	}

	ll.Pop()
	if ll.root != nil {
		t.Fatalf("LL root is not nil")
	}
}
