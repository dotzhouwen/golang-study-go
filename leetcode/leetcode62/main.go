package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)

	sort.Ints(candidates)
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
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}
		*path = append(*path, candidates[i])
		backtracking(candidates, target, i+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	sum2 := combinationSum2(candidates, target)
	fmt.Println(sum2)
}
