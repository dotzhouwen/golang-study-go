package main

import (
	"fmt"
	"math"
)

/**
滑动窗口，长度最小的子数组
*/
func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0

	minLen := math.MaxInt
	sum := 0
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			length := right - left + 1
			if length < minLen {
				minLen = length
			}
			sum -= nums[left]
			left++
		}
		right++
	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}

func arrSum(nums []int, left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += nums[i]
	}
	return sum
}

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	arrayLen := minSubArrayLen(target, nums)
	fmt.Println(arrayLen)
}
