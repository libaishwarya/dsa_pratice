package main

import "fmt"

type node struct {
	value int
	next  *node
}

type linkedlist struct {
	head *node
}

func main() {
	l1 := &linkedlist{}
	l1.add(1)
	l1.add(4)
	l1.add(2)
	l1.add(3)
	l1.add(5)

	l1.print()
	l2 := &linkedlist{}

	l2.add(13)
	l2.add(41)
	l2.add(2)
	l2.add(3)
	l1.print()

	result := mergedList(l1, l2)
	fmt.Println("merged list")
	result.print()

}

func (ll *linkedlist) add(val int) {
	if ll.head == nil {
		ll.head = &node{value: val}
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = &node{value: val}

}

func (ll *linkedlist) print() {

	if ll.head == nil {
		fmt.Println("Empty")
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
}

func mergedList(l1, l2 *linkedlist) *linkedlist {
	mergeList := &linkedlist{}
	dummy := &node{}
	current := dummy

	node1 := l1.head
	node2 := l2.head

	for node1 != nil && node2 != nil {
		if node1.value < node2.value {
			current.next = node1
			node1 = node1.next
		} else {
			current.next = node2
			node2 = node2.next
		}
		current = current.next
	}
	if node1 != nil {
		current.next = node1
	} else if node2 != nil {
		current.next = node2
	}
	mergeList.head = dummy.next

	return mergeList
}
