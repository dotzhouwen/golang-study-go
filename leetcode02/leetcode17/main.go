package main

import (
	"fmt"
)

/**
岛屿的最大面积
*/
func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, m)
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}

	maxArea := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				area := new(int)
				dfs(&grid, i, j, &visited, area)
				if *area > maxArea {
					maxArea = *area
				}
			}
		}
	}
	return maxArea
}

func dfs(grid *[][]int, i int, j int, visited *[][]bool, area *int) {
	if i >= len(*grid) || i < 0 || j >= len((*grid)[0]) || j < 0 {
		return
	}
	if (*visited)[i][j] {
		return
	}
	if (*grid)[i][j] == 0 {
		return
	}

	(*visited)[i][j] = true
	*area = *area + 1

	dfs(grid, i, j+1, visited, area)
	dfs(grid, i, j-1, visited, area)
	dfs(grid, i-1, j, visited, area)
	dfs(grid, i+1, j, visited, area)
}

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	island := maxAreaOfIsland(grid)
	fmt.Println(island)
}
