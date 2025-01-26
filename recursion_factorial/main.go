package main

import "fmt"

func main() {
	n := 10
	result := fact(n)
	fmt.Println(result)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}

	res := n * fact(n-1)
	return res
}
