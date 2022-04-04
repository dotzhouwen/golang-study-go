package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				if nums[i]+nums[j] == target {
					return []int{i, j}
				}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, exists := hash[target-nums[i]]; exists {
			return []int{hash[target-nums[i]], i}
		}
		hash[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	sum := twoSum2(nums, target)

	fmt.Println(sum)
}
