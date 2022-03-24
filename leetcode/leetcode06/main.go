package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	largest := findKthLargest(nums, 2)
	fmt.Println(largest)
}

func findKthLargest(nums []int, k int) int {
	for i := 0; i < k; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums[len(nums)-k]
}

func sortColors(nums []int) {
	start := -1
	end := len(nums)
	for fastIndex := 0; fastIndex < end; fastIndex++ {
		if nums[fastIndex] == 0 {
			start += 1
			nums[start], nums[fastIndex] = nums[fastIndex], nums[start]
		} else if nums[fastIndex] == 2 {
			end -= 1
			nums[end], nums[fastIndex] = nums[fastIndex], nums[end]
		}
	}
}
