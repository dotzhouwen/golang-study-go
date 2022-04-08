package main

import "fmt"

/**
飞地的数量 dfs
*/
func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				if grid[i][j] == 1 {
					dfs(&grid, i, j)
				}
			}
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func dfs(grid *[][]int, i int, j int) {
	if i < 0 || i >= len(*grid) || j < 0 || j >= len((*grid)[0]) || (*grid)[i][j] == 0 {
		return
	}
	(*grid)[i][j] = 0
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func main() {
	grid := [][]int{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 0, 1},
		{0, 0, 0, 0},
	}
	enclaves := numEnclaves(grid)
	fmt.Println(enclaves)
}
