package main

import (
	"fmt"
)

/**
454. 四数相加 II
*/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	hash := make(map[int]int, 0)
	for _, i := range nums1 {
		for _, j := range nums2 {
			temp := i + j
			hash[temp]++
		}
	}

	res := 0
	for _, i := range nums3 {
		for _, j := range nums4 {
			temp := i + j
			if _, exists := hash[0-temp]; exists {
				res += hash[0-temp]
			}
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}

	count := fourSumCount(nums1, nums2, nums3, nums4)
	fmt.Println(count)
}
