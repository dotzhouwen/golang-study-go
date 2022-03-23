package main

import "fmt"

/**
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 最多出现两次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。 不需要考虑数组中超出新长度后面的元素。

*/

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	ret := removeDuplicates2(nums)
	fmt.Println("ret:", ret)
	fmt.Println(nums)
}

func removeDuplicates2(nums []int) int {
	i := 1
	j := 1
	count := 1
	for ; j < len(nums); j++ {
		if nums[j] != nums[j-1] {
			count = 1
			nums[i] = nums[j]
			i++
		} else {
			count++
		}
		if count == 2 {
			nums[i] = nums[j]
			i++
		}
	}
	if len(nums) == 0 {
		return 0
	} else {
		return i
	}
}

func removeDuplicates(nums []int) int {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if countVal(nums, slowIndex, nums[fastIndex]) < 2 {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}

func countVal(nums []int, to int, val int) int {
	count := 0
	for i := 0; i < to; i++ {
		if nums[i] == val {
			count += 1
		}
	}
	return count
}
