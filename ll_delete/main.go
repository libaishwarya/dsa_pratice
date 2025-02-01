package main

import "fmt"

type node struct {
	value int
	next  *node
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	head := linkedList(list)
	head2 := linkedList(list)
	head3 := linkedList(list)

	printLinkedlist(head)
	fmt.Println("\n")
	deleteNode := deleteNode(head, 4)
	printLinkedlist(deleteNode)
	deleteFirstNode := deleteFirstNode(head2)
	fmt.Println("\n")

	printLinkedlist(deleteFirstNode)
	deleteLastnode := deleteLastnode(head3)
	fmt.Println("\n")

	printLinkedlist(deleteLastnode)

}

func deleteLastnode(head *node) *node {
	if head == nil {
		return nil
	}

	if head.next == nil {
		return nil
	}
	current := head
	var prev *node
	for current.next != nil {
		prev = current
		current = current.next

	}
	prev.next = nil
	return head
}

func deleteFirstNode(head *node) *node {
	firstNode := head.next
	return firstNode

}

func deleteNode(head *node, val int) *node {
	if head == nil {
		fmt.Println("Linkedlist is empty and no node to be deleted")
	}
	current := head
	// prev := head
	var prev *node

	for current.next != nil {

		if current.value == val {
			prev.next = current.next
			break
		}
		prev = current
		current = current.next
	}
	return head

}

func printLinkedlist(head *node) *node {
	if head == nil {
		return head
	}
	i := 0
	for head != nil {
		if i == 0 {
			fmt.Print(head.value)
		} else {
			fmt.Print("->", head.value, "")
		}
		i++
		head = head.next
	}
	return head
}

func linkedList(list []int) *node {
	var head *node
	if head == nil {
		head = &node{value: list[0]}
	}
	current := head
	for i := 1; i < len(list); i++ {
		current.next = &node{value: list[i]}
		current = current.next
	}

	current.next = nil
	return head

}
