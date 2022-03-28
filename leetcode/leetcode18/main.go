package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	duplicate := containsDuplicate(nums)
	fmt.Println(duplicate)
}

func containsDuplicate(nums []int) bool {
	containsMap := make(map[int]struct{})
	length := len(nums)
	for i := 0; i < length; i++ {
		_, exists := containsMap[nums[i]]
		if exists {
			return true
		} else {
			containsMap[nums[i]] = struct{}{}
		}
	}
	return false
}
