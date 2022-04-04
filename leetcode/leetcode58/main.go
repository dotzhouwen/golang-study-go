package main

import "fmt"

/**
leetcode 216 组合总和
*/
func combinationSum3(k int, n int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	dfs(k, n, 1, &path, &res)
	return res
}

func dfs(k int, n int, begin int, path *[]int, res *[][]int) {
	if len(*path) == k {
		if n == 0 {
			p := make([]int, k)
			copy(p, *path)
			*res = append(*res, p)
		}
		return
	}

	for i := begin; i <= 9; i++ {
		*path = append(*path, i)
		dfs(k, n-i, i+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	k := 3
	n := 9
	sum3 := combinationSum3(k, n)
	fmt.Println(sum3)
}
