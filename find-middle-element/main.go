package main

import "fmt"

type node struct {
	value int
	next  *node
}

func main() {
	node1 := &node{value: 1}
	node2 := &node{value: 2}
	node3 := &node{value: 3}
	node4 := &node{value: 4}
	node5 := &node{value: 5}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = nil

	result := middleElement(node1)
	print(result)

}

func middleElement(head *node) *node {

	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func print(head *node) {
	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}
}
