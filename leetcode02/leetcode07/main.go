package main

import "fmt"

/**
可包含重复数字的序列
*/
func permuteUnique(nums []int) [][]int {
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
	uset := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, exists := uset[nums[i]]; exists {
			continue
		}
		if !used[i] {
			uset[nums[i]] = struct{}{}
			*path = append(*path, nums[i])
			used[i] = true
			backtracking(nums, used, path, res)
			used[i] = false
			*path = (*path)[:len(*path)-1]
		}
	}
}

/**
输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
*/
func main() {
	nums := []int{1, 1, 2}
	i := permuteUnique(nums)
	fmt.Println(i)
}
