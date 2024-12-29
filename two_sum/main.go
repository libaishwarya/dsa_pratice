package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		comp := target - num
		index, found := hashMap[comp]
		if found {
			return []int{index, i}
		}
		hashMap[num] = i
	}
	return nil

}

func main() {
	nums := []int{2, 7, 11, 5}
	target := 9
	result := twoSum(nums, target)
	if result != nil {
		fmt.Printf("Indices: %v\n", result)
	} else {
		fmt.Println("No solution found.")
	}
}
