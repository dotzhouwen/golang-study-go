package main

import "fmt"

func longestOnes(nums []int, k int) int {
	left := 0
	zeroCnt := 0
	maxLen := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCnt++
		}
		for zeroCnt > k {
			if nums[left] == 0 {
				zeroCnt--
			}
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
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
