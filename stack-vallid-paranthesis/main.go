package main

import "fmt"

func main() {
	tests := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
	}
	for _, va := range tests {
		fmt.Printf("Input:%s ,output:%v\n", va, isvalid(va))
	}
}

func isvalid(test_char string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, val := range test_char {
		if opening, exists := bracketMap[val]; exists {

			if len(stack) > 0 && stack[len(stack)-1] == opening {

				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		} else {
			stack = append(stack, val)
		}
		if len(stack) == 0 {
			return true
		}
	}
	return false

}

// package main

// import (
// 	"fmt"
// )

// func isValid(s string) bool {
// 	// Create a stack to keep track of opening brackets
// 	stack := []rune{}

// 	// Map to match closing brackets with their corresponding opening brackets
// 	bracketMap := map[rune]rune{
// 		')': '(',
// 		'}': '{',
// 		']': '[',
// 	}

// 	// Iterate over the string
// 	for _, char := range s {
// 		// If the character is a closing bracket
// 		if opening, exists := bracketMap[char]; exists {
// 			// Check if the stack is not empty and the top of the stack matches the opening bracket
// 			if len(stack) > 0 && stack[len(stack)-1] == opening {
// 				// Pop the top of the stack
// 				stack = stack[:len(stack)-1]
// 			} else {
// 				// Invalid if no matching opening bracket is found
// 				return false
// 			}
// 		} else {
// 			// If it's an opening bracket, push it onto the stack
// 			stack = append(stack, char)
// 		}
// 	}

// 	// Return true if the stack is empty, meaning all brackets were matched
// 	return len(stack) == 0
// }

// func main() {
// 	// Test cases
// 	testCases := []string{
// 		"()",
// 		"()[]{}",
// 		"(]",
// 		"([)]",
// 		"{[]}",
// 	}

// 	for _, testCase := range testCases {
// 		fmt.Printf("Input: %s, Output: %v\n", testCase, isValid(testCase))
// 	}
// }
