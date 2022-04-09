package main

import "fmt"

func quickSort(nums []int, leftIndex int, rightIndex int) {
	if leftIndex >= rightIndex {
		return
	}

	left := leftIndex
	right := rightIndex
	pivot := nums[left]

	for left < right {
		for right > left && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]

		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}

	nums[left] = pivot
	quickSort(nums, leftIndex, left-1)
	quickSort(nums, right+1, rightIndex)
}

func main() {
	nums := []int{9, 8, 2, 3, 1, 4, 6, 5, 7}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
