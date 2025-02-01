package main

import "fmt"

func main() {
	r := "aa"
	m := "aab"
	result := isvalid(r, m)
	fmt.Println(result)
}

func isvalid(r, m string) bool {
	hashm := make(map[rune]int)
	for _, val := range m {
		hashm[val]++
	}
	for _, val := range r {
		if hashm[val] > 0 {
			hashm[val]--
		} else {
			return false
		}
	}
	return true

}
