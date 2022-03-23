package main

import "fmt"

/**
给定一个包含红色、白色和蓝色、共n 个元素的数组nums，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。

必须在不使用库的sort函数的情况下解决这个问题。

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]

*/

func sortColors(nums []int) {
	for i := 1; i < len(nums); i++ {
		v := nums[i]

		var t int
		for t := 0; v < nums[t] && t < i; t++ {
		}
		if t != i {
			for s := i - 1; s > t; s-- {
				nums[s] = nums[s-1]
			}
			nums[t] = v
		}
	}
}

func main() {
	nums := []int{1, 8, 9, 8, 7, 6, 5, 12}
	sortColors(nums)

	fmt.Println(nums)
}
