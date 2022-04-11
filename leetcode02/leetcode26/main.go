package main

import "fmt"

/**
删掉一个元素以后全为 1 的最长子数组
*/
func longestSubarray(nums []int) int {
	cnt := 0
	left := 0
	maxLen := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			cnt++
		}
		for cnt > 1 {
			if nums[left] == 0 {
				cnt--
			}
			left++
		}
		length := right - left
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}

func main() {
	nums := []int{1, 1, 0, 1}
	subarray := longestSubarray(nums)
	fmt.Println(subarray)
}
