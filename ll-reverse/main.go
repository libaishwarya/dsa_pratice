package main

import "fmt"

type node struct {
	value int
	next  *node
}

func main() {
	nums := createll([]int{1, 2, 3, 4, 5})
	// nums := createll([]int{})

	// printRecursion(nums)
	printRevRecursion(nums)
	// revers := reverse(nums)
	// printlist(revers)

}

func createll(num []int) *node {
	if len(num) == 0 {
		return nil
	}
	head := &node{value: num[0]}
	current := head
	for i := 1; i < len(num); i++ {
		current.next = &node{value: num[i]}
		current = current.next
	}
	return head
}

func reverse(head *node) *node {
	current := head
	var prev *node

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	return prev
}

func printlist(head *node) {
	current := head

	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func printRecursion(n *node) {
	if n == nil {
		return
	}

	fmt.Println(n.value)
	printRecursion(n.next)
}

func printRevRecursion(n *node) {
	if n == nil {
		return
	}

	printRevRecursion(n.next)
	fmt.Println(n.value)
}
