package main

import "fmt"

/**
回溯 子集 解法2
*/
func subsets(nums []int) [][]int {
	var path []int
	var res [][]int
	backtracking(nums, 0, &path, &res)

	return res
}

func backtracking(nums []int, startIndex int, path *[]int, res *[][]int) {
	p := make([]int, len(*path))
	copy(p, *path)
	*res = append(*res, p)

	for i := startIndex; i < len(nums); i++ {
		*path = append(*path, nums[i])
		backtracking(nums, i+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	i := subsets(nums)
	fmt.Println(i)
}
