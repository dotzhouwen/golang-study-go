package main

import "fmt"

func maxProfit(prices []int) int {
	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}

func main() {
	prices := []int{7, 6, 4, 3, 1}
	profit := maxProfit(prices)
	fmt.Println(profit)
}
