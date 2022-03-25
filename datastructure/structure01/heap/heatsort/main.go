package main

import "fmt"

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

func buildHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, len(nums))
	}
}

func heapSort(nums []int) {
	buildHeap(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		adjustHeap(nums, 0, i)
	}
}

func main() {
	num := []int{98, 48, 777, 63, 57, 433, 23, 1112, 1, 0}
	fmt.Println(num)
	heapSort(num)
	fmt.Println(num)
}
