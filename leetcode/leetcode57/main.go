package main

import "fmt"

/**
回溯算法 组合问题 2
*/
func combine(n int, k int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	dfs(n, k, 1, &path, &res)
	return res
}

func dfs(n int, k int, begin int, path *[]int, res *[][]int) {
	if len(*path) == k {
		p := make([]int, k)
		copy(p, *path)
		*res = append(*res, p)
	}

	for i := begin; i <= n; i++ {
		*path = append(*path, i)
		dfs(n, k, i+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	res := combine(4, 3)
	fmt.Println(res)
}
