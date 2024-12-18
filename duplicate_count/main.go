package main

import "fmt"

func main() {
	arr := []int{1, 2, 7, 4, 2, 1}
	duplicates := duplicate(arr)
	if len(duplicates) > 0 {
		fmt.Println(duplicates)
	} else {
		fmt.Println("No duplicates found")
	}

}

func duplicate(arr []int) []int {
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num] = freq[num] + 1
	}
	var dup []int
	for num, count := range freq {
		if count > 1 {
			dup = append(dup, num)
		}
	}
	return dup
}
