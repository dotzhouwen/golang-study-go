package main

import "fmt"

/**
移除排序数组的重复项
*/
func removeDuplicates2(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slow := 2
	fast := 2
	for fast < len(nums) {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func removeDuplicates(nums []int) int {
	hash := make(map[int]int, 0)
	slow := 0
	fast := 0

	for fast < len(nums) {
		if _, exists := hash[nums[fast]]; exists {
			count := hash[nums[fast]]
			if count < 2 {
				nums[slow] = nums[fast]
				hash[nums[fast]]++
				slow++
			}
		} else {
			nums[slow] = nums[fast]
			hash[nums[fast]] = 1
			slow++
		}
		fast++
	}
	return slow
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	duplicates := removeDuplicates(nums)
	fmt.Println(duplicates)

	fmt.Println(nums)
}
