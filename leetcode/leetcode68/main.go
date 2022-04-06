package main

import (
	"fmt"
	"math"
)

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))

	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		dp[i] = int(math.Min(float64(dp[i-1]), float64(dp[i-2]))) + cost[i]
	}
	return int(math.Min(float64(dp[len(cost)-1]), float64(dp[len(cost)-2])))
}

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	stairs := minCostClimbingStairs(cost)
	fmt.Println(stairs)
}
