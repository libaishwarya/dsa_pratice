package main

import "fmt"

func majority(num []int) int {
	hashmap := make(map[int]int)
	for _, val := range num {
		hashmap[val] = hashmap[val] + 1
		if hashmap[val] > len(num)/2 {
			return val
		}
	}
	return -1
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majority(nums))

	num1 := []int{3, 1, 3, 1, 1}
	fmt.Println(majority(num1))
}
