package main

import "fmt"

/**
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:

输入: nums = [0]
输出: [0]
*/
func moveZeroes(nums []int) {
	var p = 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != 0 {
			nums[p] = nums[fastIndex]
			p++
		}
	}
	for i := len(nums) - 1; i >= p; i-- {
		nums[i] = 0
	}
}

func moveZeros(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp

			i++
		}
	}
}

func moveZeros2(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] == 0 {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
}

func main() {
	nums := []int{1, 9, 8, 0, 0, 7, 6, 0, 5}
	moveZeros2(nums)
	fmt.Println(nums)
}
