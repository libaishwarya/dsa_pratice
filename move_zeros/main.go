package main

import "fmt"

func moveZeros(nums []int) {
	track := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			fmt.Println(nums[i], nums[track])
			nums[i], nums[track] = nums[track], nums[i]
			track++
			fmt.Println(nums)
		}
	}
}

func main() {
	nums := []int{1, 0, 5, 6, 0, 10, 12}
	moveZeros(nums)
	fmt.Println(nums)
}
