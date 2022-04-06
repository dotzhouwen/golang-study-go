package main

import "fmt"

/**
78. 子集
*/

func subsets(nums []int) [][]int {
	var path []int
	var res [][]int

	for i := 0; i <= len(nums); i++ {
		backtracking(nums, i, 0, &path, &res)
	}
	return res
}

func backtracking(nums []int, count int, startIndex int, path *[]int, res *[][]int) {
	if len(*path) >= count {
		p := make([]int, len(*path))
		copy(p, *path)
		*res = append(*res, p)
		return
	}

	for i := startIndex; i < len(nums); i++ {
		*path = append(*path, nums[i])
		backtracking(nums, count, i+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	i := subsets(nums)
	fmt.Println(i)
}
