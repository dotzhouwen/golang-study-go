package main

import (
	"fmt"
	"math"
)

/**
1695. 删除子数组的最大得分, 滑动窗口
*/
func maximumUniqueSubarray(nums []int) int {
	hash := make(map[int]int, 0)
	maxSum := math.MinInt64

	sum := 0
	j := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if _, exists := hash[v]; exists {
			hash[v]++
		} else {
			hash[v] = 1
		}

		count, _ := hash[v]
		for count > 1 {
			hash[nums[j]]--
			sum -= nums[j]
			j++
			count, _ = hash[v]
		}

		sum += v
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func main() {
	nums := []int{5, 2, 1, 2, 5, 2, 1, 2, 5}
	subarray := maximumUniqueSubarray(nums)
	fmt.Println(subarray)
}
