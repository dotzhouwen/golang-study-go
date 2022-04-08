package main

import "fmt"

/**
统计封闭岛屿的数目dfs
*/
func closedIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, m)
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}

	totalCount := 0
	boundCount := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && !visited[i][j] {
				isBound := new(bool)
				dfs(&grid, i, j, &visited, isBound)
				totalCount++

				if *isBound {
					boundCount++
				}
			}
		}
	}

	return totalCount - boundCount
}

func dfs(grid *[][]int, i int, j int, visited *[][]bool, isBound *bool) {
	if i < 0 || i >= len(*grid) || j < 0 || j >= len((*grid)[0]) {
		*isBound = true
		return
	}

	if (*grid)[i][j] == 1 {
		return
	}

	if (*visited)[i][j] {
		return
	}

	(*visited)[i][j] = true
	dfs(grid, i+1, j, visited, isBound)
	dfs(grid, i-1, j, visited, isBound)
	dfs(grid, i, j+1, visited, isBound)
	dfs(grid, i, j-1, visited, isBound)
}

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0},
	}
	island := closedIsland(grid)
	fmt.Println(island)
}
