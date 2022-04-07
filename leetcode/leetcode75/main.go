package main

import (
	"fmt"
)

/**
最大岛屿面积 bfs

*/
func maxAreaOfIsland(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i, _ := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				area := bfs(&grid, i, j, &visited)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func bfs(grid *[][]int, i int, j int, visited *[][]bool) int {
	queue := make([][]int, 0)
	queue = append(queue, []int{i, j})

	area := 0
	for len(queue) > 0 {
		e := queue[0]
		x, y := e[0], e[1]
		queue = queue[1:]

		for x >= 0 && x < len(*grid) && y >= 0 && y < len((*grid)[0]) && (*grid)[x][y] == 1 && !(*visited)[x][y] {
			(*visited)[x][y] = true
			area += 1

			queue = append(queue, []int{x + 1, y})
			queue = append(queue, []int{x - 1, y})
			queue = append(queue, []int{x, y + 1})
			queue = append(queue, []int{x, y - 1})
		}
	}
	return area
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
