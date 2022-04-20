package main

import "fmt"

/**
1034. 边界着色

*/
func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	originColor := grid[row][col]

	border := make([][]int, len(grid))
	for i, _ := range border {
		border[i] = make([]int, len(grid[0]))
	}

	dfs(&grid, row, col, originColor, &border)

	for i := 0; i < len(border); i++ {
		for j := 0; j < len(border[0]); j++ {
			if border[i][j] > 1 {
				grid[i][j] = color
			}
		}
	}
	return grid
}

func dfs(grid *[][]int, i int, j int, originColor int, border *[][]int) {
	if i < 0 || i >= len(*grid) || j < 0 || j >= len((*grid)[0]) {
		return
	}
	if (*grid)[i][j] != originColor {
		return
	}
	if (*border)[i][j] > 0 {
		return
	}

	(*border)[i][j]++
	if i == 0 || i == len(*grid)-1 || j == 0 || j == len((*grid)[0])-1 {
		(*border)[i][j]++
	}
	if i >= 1 && (*grid)[i-1][j] != originColor {
		(*border)[i][j]++
	}
	if i < len(*grid)-1 && (*grid)[i+1][j] != originColor {
		(*border)[i][j]++
	}
	if j >= 1 && (*grid)[i][j-1] != originColor {
		(*border)[i][j]++
	}
	if j < len((*grid)[0])-1 && (*grid)[i][j+1] != originColor {
		(*border)[i][j]++
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, direction := range directions {
		dfs(grid, i+direction[0], j+direction[1], originColor, border)
	}
}

func main() {
	grid := [][]int{
		{1, 1},
		{1, 2},
	}
	row := 0
	col := 0
	color := 3

	border := colorBorder(grid, row, col, color)
	for _, row := range border {
		fmt.Println(row)
	}
}
