package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		for left < right && !alphanum(rune(s[left])) {
			left++
		}
		for left < right && !alphanum(rune(s[right])) {
			right--
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		right--
		left++
	}
	return true

}
func alphanum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
