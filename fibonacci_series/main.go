package main

// iterative method
//it uses a for loop to calculate each Fibonacci number based on the previous two numbers
import "fmt"

func main() {
	n := 10
	result := fibo(n)
	fmt.Println(result)
}

func fibo(num int) []int {
	fiboSeries := []int{0, 1}
	for i := 2; i < num; i++ {
		next := fiboSeries[i-1] + fiboSeries[i-2]
		fiboSeries = append(fiboSeries, next)
	}
	return fiboSeries
}
