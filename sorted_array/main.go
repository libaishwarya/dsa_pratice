package main

import "fmt"

func sortArray(n []int) []int {
	left := 0
	right := len(n) - 1
	pos := len(n) - 1
	result := make([]int, len(n))

	for left <= right {
		leftSq := n[left] * n[left]
		rightSq := n[right] * n[right]
		if leftSq < rightSq {
			result[pos] = rightSq
			right--

		} else {
			result[pos] = leftSq
			left++
		}
		pos--

	}
	return result
}

func main() {
	n := []int{-3, 1, 5, 10}
	fmt.Println(sortArray(n))
}
