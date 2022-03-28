package main

import "fmt"

/**
旋转数组
1, 2, 3, 4, 5, 6, 7
7, 6, 5, 4, 3, 2, 1
5, 6, 7, 1, 2, 3, 4
*/

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums))
	reverse(nums, 0, k)
	reverse(nums, k, len(nums))
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end-1] = nums[end-1], nums[start]
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 1
	rotate(nums, k)

	fmt.Println(nums)
}
