package main

import "fmt"

/**
循环不变量

*/
func removeDuplicates(nums []int) int {
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

func main() {
	nums := []int{1, 1, 2}
	duplicates := removeDuplicates(nums)

	fmt.Println(nums)
	fmt.Println(duplicates)
}
