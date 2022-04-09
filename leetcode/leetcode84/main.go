package main

import "fmt"

/**
移除元素
*/
func removeElement(nums []int, val int) int {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	element := removeElement(nums, val)
	fmt.Println(element)
	fmt.Println(nums)
}
