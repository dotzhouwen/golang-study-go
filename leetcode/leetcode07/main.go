package main

/**
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
*/

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	largest := findKthLargest(nums, 1)
	fmt.Println(largest)
}

func findKthLargest(nums []int, k int) int {

	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, len(nums))
	}

	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		adjustHeap(nums, 0, i)
	}
	return nums[0]
}

func adjustHeap(nums []int, pos int, length int) {
	for {
		child := pos*2 + 1
		if child > length-1 {
			break
		}
		if child+1 <= length-1 && nums[child+1] > nums[child] {
			child += 1
		}
		if nums[pos] < nums[child] {
			nums[pos], nums[child] = nums[child], nums[pos]
			pos = child
		} else {
			break
		}
	}
}
