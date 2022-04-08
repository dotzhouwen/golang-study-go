package main

import "fmt"

/*
被围绕的区域
*/
func solve(board [][]byte) {

	visited := make([][]bool, len(board))
	for i, _ := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1 {
				if board[i][j] == 'O' {
					dfs(&board, i, j, &visited)
				}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(grid *[][]byte, i int, j int, visited *[][]bool) {
	if i < 0 || i >= len(*grid) || j < 0 || j >= len((*grid)[0]) {
		return
	}
	if (*grid)[i][j] != 'O' {
		return
	}
	if (*visited)[i][j] {
		return
	}
	(*visited)[i][j] = true
	dfs(grid, i-1, j, visited)
	dfs(grid, i+1, j, visited)
	dfs(grid, i, j-1, visited)
	dfs(grid, i, j+1, visited)
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	solve(board)

	fmt.Println(board)
}
