package main

import (
	"fmt"
)

/**
79. 单词搜索
*/
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	chars := []byte(word)

	visited := make([][]bool, m)
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(&board, &visited, i, j, chars, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board *[][]byte, visited *[][]bool, i int, j int, chars []byte, curIndex int) bool {
	if curIndex >= len(chars) {
		return true
	}
	if i < 0 || i >= len(*board) || j < 0 || j >= len((*board)[0]) {
		return false
	}
	if chars[curIndex] != (*board)[i][j] {
		return false
	}
	if (*visited)[i][j] {
		return false
	}
	(*visited)[i][j] = true

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, direction := range directions {
		b := dfs(board, visited, i+direction[0], j+direction[1], chars, curIndex+1)
		if b {
			return true
		}
	}
	(*visited)[i][j] = false
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCES"
	b := exist(board, word)
	fmt.Println(b)
}
