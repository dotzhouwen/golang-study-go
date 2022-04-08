package main

import "fmt"

/**
岛屿数量bfs
*/
func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i, _ := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				bfs(&grid, i, j, &visited)
				count++
			}
		}
	}
	return count
}

func bfs(grid *[][]byte, i int, j int, visited *[][]bool) {
	queue := make([][]int, 0)
	queue = append(queue, []int{i, j})

	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]
		x, y := e[0], e[1]

		if x >= 0 && x < len(*grid) && y >= 0 && y < len((*grid)[0]) && (*grid)[x][y] == '1' && !(*visited)[x][y] {
			(*visited)[x][y] = true

			queue = append(queue, []int{x + 1, y})
			queue = append(queue, []int{x - 1, y})
			queue = append(queue, []int{x, y + 1})
			queue = append(queue, []int{x, y - 1})
		}
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '1'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	islands := numIslands(grid)
	fmt.Println(islands)
}
