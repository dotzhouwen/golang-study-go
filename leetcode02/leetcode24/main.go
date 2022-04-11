package main

import "fmt"

/**
连续1的个数
*/
func findMaxConsecutiveOnes(nums []int) int {
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
		length := right - left + 1
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}

func main() {
	nums := []int{1, 0, 1, 1, 0}
	ones := findMaxConsecutiveOnes(nums)
	fmt.Println(ones)
}
