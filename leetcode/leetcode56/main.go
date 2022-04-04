package main

import "fmt"

func combine(n int, k int) [][]int {

	used := make([]bool, n+1)
	res := make([][]int, 0)
	path := make([]int, 0)

	dfs(n, k, &path, &res, used, -1)
	return res
}

/**
回溯算法 解决 leetcode 77 组合问题
*/
func dfs(n int, k int, path *[]int, res *[][]int, used []bool, cur int) {
	if len(*path) == k {
		p := make([]int, k)
		copy(p, *path)
		*res = append(*res, p)
		return
	}

	for i := 1; i <= n; i++ {
		if !used[i] && i > cur {
			*path = append(*path, i)
			used[i] = true

			dfs(n, k, path, res, used, i)
			used[i] = false
			*path = (*path)[:len(*path)-1]
		}
	}
}

func main() {
	i := combine(4, 4)
	fmt.Println(i)
}
