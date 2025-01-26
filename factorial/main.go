package main

import "fmt"

func main() {
	n := 10
	result := factorial(n)
	fmt.Println(result)
}

func factorial(n int) int {
	res := n * (n - 1)

	for i := 2; i < n; i++ {
		res = res * (n - i)
	}
	return res
}
