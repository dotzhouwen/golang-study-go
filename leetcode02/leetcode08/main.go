package main

import "fmt"

func solveNQueens(n int) [][]string {
	var path [][]int
	var queens [][]string
	backtracking(n, 0, &path, &queens)

	return queens
}

/**
输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
*/
func backtracking(n int, level int, path *[][]int, queens *[][]string) {
	if level >= n {
		lines := make([]string, 0)
		for _, loc := range *path {
			strLoc := make([]rune, n)
			for i := 0; i < n; i++ {
				strLoc[i] = '.'
			}
			strLoc[loc[1]] = 'Q'
			lines = append(lines, string(strLoc))
		}
		*queens = append(*queens, lines)
		return
	}
	for i := 0; i < n; i++ {
		if canLocate(level, i, path) {
			*path = append(*path, []int{level, i})
			backtracking(n, level+1, path, queens)
			*path = (*path)[:len(*path)-1]
		}
	}
}

func canLocate(i int, j int, path *[][]int) bool {
	for _, loc := range *path {
		li := loc[0]
		lj := loc[1]

		if j == lj {
			return false
		}
		if i-j == li-lj {
			return false
		}
		if i+j == li+lj {
			return false
		}
	}
	return true
}

func main() {
	queens := solveNQueens(8)
	for _, queue := range queens {
		fmt.Println(queue)
	}
}
