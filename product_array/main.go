package main

import "fmt"

func productExceptSelf(nums []int) []int {
	result := []int{}
	for i := range nums {
		acc := 1
		for j, v2 := range nums {
			if j == i {
				continue
			}
			acc = acc * v2
		}
		result = append(result, acc)
	}
	return result
}

func main() {
	input := []int{1, 2, 3, 4}

	fmt.Println(productExceptSelf(input))

}
