package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	for j := 0; j < k; j++ {
		for i := 0; i < len(nums)-j-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	return nums[len(nums)-k]
}

func main() {
	nums := []int{6, 3, 2, 3, 1, 2, 4, 5, 5}
	k := 4
	largest := findKthLargest(nums, k)
	fmt.Println(largest)
}
