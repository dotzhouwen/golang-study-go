package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	backtracking(&candidates, target, &path, &res, 0, 0)
	return res
}

/**
组合问题 回溯法
*/
func backtracking(candidates *[]int, target int, path *[]int, res *[][]int, cur int, sum int) {
	if sum >= target {
		if sum == target {
			p := make([]int, len(*path))
			copy(p, *path)
			*res = append(*res, p)
		}
		return
	}
	for i := cur; i < len(*candidates); i++ {
		*path = append(*path, (*candidates)[i])
		sum += (*candidates)[i]
		backtracking(candidates, target, path, res, i, sum)
		*path = (*path)[:len(*path)-1]
		sum -= (*candidates)[i]
	}
}

func main() {
	candidates := []int{2, 3, 5}
	target := 8
	sum := combinationSum(candidates, target)
	fmt.Println(sum)
}
