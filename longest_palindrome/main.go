package main

import (
	"fmt"
)

func longestPalindrome(s string) int {
	trackMap := make(map[rune]int)
	for _, val := range s {
		trackMap[val]++
	}
	hasOdd := false
	length := 0
	for _, v1 := range trackMap {
		if v1%2 == 0 {
			length = length + v1
		} else {
			length = length + v1 - 1
			hasOdd = true
		}
	}
	if hasOdd {
		length = length + 1
	}
	return length
}
func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}
