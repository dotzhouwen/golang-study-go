package main

import "fmt"

/**
最大连续 1 的个数 III
*/
func longestOnes(nums []int, k int) int {
	left := 0
	cnt := 0
	maxLen := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			cnt++
		}
		for cnt > k {
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
	nums := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	K := 3
	ones := longestOnes(nums, K)
	fmt.Println(ones)
}
