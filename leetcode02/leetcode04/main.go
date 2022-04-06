package main

import (
	"fmt"
	"sort"
)

/*
子集2 去重
*/
func subsetsWithDup(nums []int) [][]int {
	var path []int
	var res [][]int
	sort.Ints(nums)
	backtracking(nums, 0, &path, &res)
	return res
}

func backtracking(nums []int, startIndex int, path *[]int, res *[][]int) {
	p := make([]int, len(*path))
	copy(p, *path)
	*res = append(*res, p)

	for i := startIndex; i < len(nums); i++ {
		if i > startIndex && nums[i] == nums[i-1] {
			continue
		}
		*path = append(*path, nums[i])
		backtracking(nums, i+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

/**
[[],[1],[1,2],[1,2,2],[2],[2,2]]
*/
func main() {
	nums := []int{1, 2, 2}
	dup := subsetsWithDup(nums)
	fmt.Println(dup)
}
