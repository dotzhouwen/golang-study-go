package main

import "fmt"

/**
不重复数字全排列
*/
func permute(nums []int) [][]int {
	var path []int
	var res [][]int
	used := make([]bool, len(nums))
	backtracking(nums, used, &path, &res)
	return res
}

func backtracking(nums []int, used []bool, path *[]int, res *[][]int) {
	if len(*path) == len(nums) {
		p := make([]int, len(*path))
		copy(p, *path)
		*res = append(*res, p)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			*path = append(*path, nums[i])
			used[i] = true
			backtracking(nums, used, path, res)
			used[i] = false
			*path = (*path)[:len(*path)-1]
		}
	}
}

/**
nums = [1,2,3]
[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
func main() {
	nums := []int{1, 1, 2}
	i := permute(nums)
	fmt.Println(i)
}
