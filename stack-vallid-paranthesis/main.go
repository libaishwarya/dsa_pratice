package main

import "fmt"

func main() {
	tests := []string{
		"(){}}{",
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

// 	for _, val := range test_char {
// 		if opening, exists := bracketMap[val]; exists {

// 			if len(stack) > 0 && stack[len(stack)-1] == opening {

// 				stack = stack[:len(stack)-1]
// 			} else {
// 				return false
// 			}

// 		} else {
// 			stack = append(stack, val)
// 		}
// 		if len(stack) == 0 {
// 			return true
// 		}
// 	}
// 	return false

// }
func isvalid(s string) bool {
	bracket := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []rune{}

	for _, val := range s {
		if openBracket, exists := bracket[val]; exists {
			// If stack is empty or top of the stack does not match expected opening bracket, return false
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				return false
			}
			// Pop the last element (valid match found)
			stack = stack[:len(stack)-1]
		} else {
			// Push opening brackets onto the stack
			stack = append(stack, val)
		}

	}
	return len(stack) == 0
}
