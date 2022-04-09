package main

import (
	"fmt"
)

/**
最小高度树
*/
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	adjTable := make([][]int, n)
	degrees := make([]int, n)

	for _, edge := range edges {
		adjTable[edge[0]] = append(adjTable[edge[0]], edge[1])
		adjTable[edge[1]] = append(adjTable[edge[1]], edge[0])
		degrees[edge[0]]++
		degrees[edge[1]]++
	}

	queue := make([]int, 0)
	for i, degree := range degrees {
		if degree == 1 {
			queue = append(queue, i)
		}
	}
	var res []int

	for len(queue) > 0 {
		size := len(queue)
		res = make([]int, 0)
		for i := 0; i < size; i++ {
			e := queue[0]
			queue = queue[1:]
			res = append(res, e)

			for _, v := range adjTable[e] {
				degrees[v]--
				if degrees[v] == 1 {
					queue = append(queue, v)
				}
			}
		}
	}
	return res
}

func main() {
	n := 6
	edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}
	trees := findMinHeightTrees(n, edges)
	fmt.Println(trees)
}
