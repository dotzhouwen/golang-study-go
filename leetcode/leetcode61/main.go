package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	backtracking(candidates, target, 0, &path, &res)
	return res
}

func backtracking(candidates []int, target int, startIndex int, path *[]int, res *[][]int) {
	sum := 0
	for _, v := range *path {
		sum += v
	}
	if sum >= target {
		if sum == target {
			p := make([]int, len(*path))
			copy(p, *path)
			*res = append(*res, p)
		}
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		*path = append(*path, candidates[i])
		backtracking(candidates, target, i, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	candidates := []int{2}
	target := 1
	sum := combinationSum(candidates, target)
	fmt.Println(sum)
}
