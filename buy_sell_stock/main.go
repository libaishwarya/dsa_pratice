package main

import "fmt"

func maxProfit(prices []int) int {
	min := prices[0]
	max_profit := 0
	for _, val := range prices {
		if val < min {
			min = val

		} else {
			profit := val - min
			if profit > max_profit {
				max_profit = profit
			}
		}
	}
	return max_profit

}
func main() {
	price := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(price))

}
