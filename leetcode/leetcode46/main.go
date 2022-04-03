package main

import "fmt"

/**
leetcode 200 岛屿数量
*/
func numIslands(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])

	visited := make([][]bool, row)
	for i, _ := range visited {
		visited[i] = make([]bool, col)
	}

	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("grid[%d][%d]=%d\n", i, j, grid[i][j])

			if grid[i][j] == '1' {

				if !visited[i][j] {
					dfs(grid, i, j, &visited)
					count++
				}
			}
		}
	}

	return count
}

func dfs(grid [][]byte, i int, j int, visited *[][]bool) {
	row := len(grid)
	col := len(grid[0])

	if i < 0 || i >= row || j < 0 || j >= col {
		return
	}

	if grid[i][j] == '0' || (*visited)[i][j] == true {
		return
	}

	(*visited)[i][j] = true
	dfs(grid, i+1, j, visited)
	dfs(grid, i-1, j, visited)
	dfs(grid, i, j+1, visited)
	dfs(grid, i, j-1, visited)
}

func main() {
	grid := [][]byte{
		{1, 1, 1, 1, 1},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	islands := numIslands(grid)
	fmt.Println(islands)
}
