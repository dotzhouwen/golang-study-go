package main

import "fmt"

/**
删除有序数组中的重复项
*/
func removeDuplicates(nums []int) int {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	duplicates := removeDuplicates(nums)
	fmt.Println(duplicates)

	fmt.Println(nums)
}
