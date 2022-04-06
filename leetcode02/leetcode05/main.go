package main

import (
	"fmt"
)

/**
491. 递增子序列
*/
func findSubsequences(nums []int) [][]int {
	var path []int
	var res [][]int

	backtracking(nums, 0, &path, &res)
	return res
}

func backtracking(nums []int, startIndex int, path *[]int, res *[][]int) {
	if len(*path) >= 2 {
		p := make([]int, len(*path))
		copy(p, *path)
		*res = append(*res, p)
	}

	uset := make(map[int]struct{})
	for i := startIndex; i < len(nums); i++ {
		if _, exists := uset[nums[i]]; exists {
			continue
		}
		uset[nums[i]] = struct{}{}
		if len(*path) > 0 {
			last := (*path)[len(*path)-1]
			if nums[i] >= last {
				*path = append(*path, nums[i])
				backtracking(nums, i+1, path, res)
				*path = (*path)[:len(*path)-1]
			}
		} else {
			*path = append(*path, nums[i])
			backtracking(nums, i+1, path, res)
			*path = (*path)[:len(*path)-1]
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 1, 1, 1, 1}
	subsequences := findSubsequences(nums)
	fmt.Println(subsequences)
}
