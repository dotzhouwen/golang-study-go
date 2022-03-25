package main

import (
	"fmt"
)

/**
给定一个包含红色、白色和蓝色、共n 个元素的数组nums，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。

必须在不使用库的sort函数的情况下解决这个问题。

输入：nums = [2,0,2,1,1,0,0,1,2]
输出：[0,0,1,1,2,2]

*/

func sortColors(nums []int) {
	redIndex := 0
	blueIndex := len(nums) - 1

	for i := 0; i < blueIndex+1; i++ {
		if nums[i] == 0 {
			tmp := nums[redIndex]
			nums[redIndex] = nums[i]
			nums[i] = tmp
			redIndex++
		} else if nums[i] == 2 {
			tmp := nums[blueIndex]
			nums[blueIndex] = nums[i]
			nums[i] = tmp
			blueIndex--
			i--
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0, 0, 1, 2}
	sortColors(nums)

	fmt.Println(nums)
}
