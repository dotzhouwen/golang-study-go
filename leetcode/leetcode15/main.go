package main

import "fmt"

/**
剑指 Offer 12. 矩阵中的路径
*/

func exist(board [][]byte, word string) bool {
	words := []byte(word)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, words, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, words []byte, i int, j int, k int) bool {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || board[i][j] != words[k] {
		return false
	}
	if k == len(words)-1 {
		return true
	}
	board[i][j] = 0
	res := dfs(board, words, i-1, j, k+1) || dfs(board, words, i+1, j, k+1) || dfs(board, words, i, j-1, k+1) || dfs(board, words, i, j+1, k+1)
	board[i][j] = words[k]
	return res
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "BCCFB"

	ret := exist(board, word)
	fmt.Println(ret)
}
