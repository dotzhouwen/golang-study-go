package main

import (
	"fmt"
	"math"
)

/**
子数组最大平均数 I 滑动窗口2
*/
func findMaxAverage2(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	i := k
	temp := sum
	for i < len(nums) {
		temp = temp + nums[i] - nums[i-k]
		sum = int(math.Max(float64(sum), float64(temp)))
		i++
	}
	return float64(sum) / float64(k)
}

/**
子数组最大平均数 I 滑动窗口
*/
func findMaxAverage(nums []int, k int) float64 {
	left := 0
	right := left + k

	var sum int
	var average float64
	for right <= len(nums) {
		if left == 0 {
			sum = numsSum(nums, left, right)
			average = float64(sum) / float64(right-left)
		} else {
			sum = sum - nums[left-1]
			sum = sum + nums[right-1]
			avg := float64(sum) / float64(right-left)
			if avg > average {
				average = avg
			}
		}
		left++
		right++
	}
	return average
}

func numsSum(nums []int, left int, right int) int {
	sum := 0
	for i := left; i < right; i++ {
		sum += nums[i]
	}
	return sum
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	average := findMaxAverage2(nums, k)
	fmt.Println(average)
}
