package main

import "fmt"

/**
485. 最大连续 1 的个数 简单
*/
func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count = 0
		} else {
			count++
			if count > maxCount {
				maxCount = count
			}
		}
	}
	return maxCount
}

func main() {
	nums := []int{1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1}
	ones := findMaxConsecutiveOnes(nums)
	fmt.Println(ones)
}
