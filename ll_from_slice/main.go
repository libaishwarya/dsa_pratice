package main

import "fmt"

type node struct {
	value int
	next  *node
}

func main() {
	num := []int{1, 2, 3, 4, 5}
	result := createList(num)
	fmt.Println(result)
	printL(result)

}

func createList(nums []int) *node {
	if len(nums) == 0 {
		fmt.Println("no numbers")
	}

	head := &node{value: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.next = &node{value: num}
		current = current.next
		fmt.Println(current)
	}
	return head
}

func printL(head *node) {
	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}

}
