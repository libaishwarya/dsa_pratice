package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func main() {
	ll := LinkedList{}

	ll.add(1)
	ll.add(2)
	ll.add(3)

	ll.print()
	fmt.Println(ll.last().value)
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) isEmpty() bool {
	return ll.head == nil
}

func (ll *LinkedList) add(i int) {
	if ll.isEmpty() {
		ll.head = &Node{value: i}
		return
	}

	lastNode := ll.last()
	lastNode.next = &Node{value: i}
}

func (ll *LinkedList) last() *Node {
	if ll.isEmpty() {
		return &Node{}
	}

	n := ll.head
	for {
		if n.next == nil {
			return n
		}
		n = n.next
	}
}

func (ll *LinkedList) print() {
	if ll.isEmpty() {
		fmt.Println("empty")
		return
	}

	n := ll.head
	for {
		fmt.Println(n.value)
		if n.next == nil {
			break
		}

		n = n.next
	}
}
