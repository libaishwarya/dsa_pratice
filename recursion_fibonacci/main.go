package main

import "fmt"

func main() {
	n := 5
	result := recursion(n)
	fmt.Println(result)
}

func recursion(num int) []int {
	if num == 0 {
		return nil
	}

	if num == 1 {
		return []int{0}
	}

	if num == 2 {
		return []int{0, 1}
	}

	fibo := recursion(num - 1)
	next := fibo[num-3] + fibo[num-2]
	fibo = append(fibo, next)
	return fibo
}
