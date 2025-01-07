package main

import "fmt"

func main() {
	s := "anagam"
	t := "nagaram"
	fmt.Println(anagram(s, t))

}

func anagram(s, t string) bool {
	hashMap := make(map[rune]int)

	for _, val := range s {
		hashMap[val]++
	}
	for _, val := range t {
		hashMap[val]--
		if hashMap[val] < 0 {
			return false
		}
	}
	for i, val := range hashMap {
		if val != 0 {
			return false
		}
		fmt.Println(i, val)

	}

	return true
}
