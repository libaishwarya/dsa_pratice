package main

import "fmt"

type node struct {
	value int
	next  *node
}

func main() {
	node1 := &node{value: 1}
	node2 := &node{value: 2}
	node3 := &node{value: 2}
	node4 := &node{value: 1}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = nil

	result := ispalindrom(node1)
	fmt.Println(result)
}

func ispalindrom(head *node) bool {
	if head == nil || head.next == nil {
		return true
	}

	slow := head
	fast := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	reverse := reverse(slow)
	first, second := head, reverse
	for second != nil {
		if first.value != second.value {
			return false
		}
		first = first.next
		second = second.next

	}
	return true
}

func reverse(head *node) *node {

	var prev *node
	current := head

	for current != nil {

		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	return prev

}
