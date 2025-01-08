package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("Common Prefix:", findCommonPrefixAmongStrings(strs))
}

func findCommonPrefixAmongStrings(strs []string) string {
	prefix := strs[0]
	for _, val := range strs[1:] {
		prefix = comparePrefix(prefix, val)
		if prefix == "" {
			break
		}
	}
	return prefix
}

func comparePrefix(str1, str2 string) string {
	minLength := len(str1)
	if minLength > len(str2) {
		minLength = len(str2)
	}
	i := 0
	for i < minLength && str1[i] == str2[i] {
		i++
	}
	return str1[:i]
}
