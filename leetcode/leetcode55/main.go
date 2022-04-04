package main

import "fmt"

func permute(nums []int) [][]int {
	used := make([]bool, len(nums))
	res := make([][]int, 0)
	path := make([]int, 0)
	dfs(nums, len(nums), used, &path, &res)
	return res
}

func dfs(nums []int, size int, used []bool, path *[]int, res *[][]int) {
	if len(*path) == size {
		p := make([]int, size)
		copy(p, *path)
		*res = append(*res, p)
	}

	for i := 0; i < size; i++ {
		if !used[i] {
			*path = append(*path, nums[i])
			used[i] = true
			dfs(nums, size, used, path, res)

			used[i] = false
			*path = (*path)[0 : len(*path)-1]
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := permute(nums)

	fmt.Println(res)
}
