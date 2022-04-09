package main

import "fmt"

/**
最长连续递增序列
*/
func findLengthOfLCIS(nums []int) int {
	count := 1
	maxCount := count
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 1
		}
	}
	return maxCount
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	i := findLengthOfLCIS(nums)
	fmt.Println(i)
}
