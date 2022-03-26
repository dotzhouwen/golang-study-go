package main

import "fmt"

func main() {
	paths := uniquePaths(3, 3)
	fmt.Println(paths)
}

/**
一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

*/
func uniquePaths(m int, n int) int {
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				f[i][j] = 1
			} else {
				fmt.Printf("f[%d][%d]=f[%d][%d]+f[%d][%d]\n", i, j, i-1, j, i, j-1)
				f[i][j] = f[i-1][j] + f[i][j-1]
				fmt.Printf("f[%d][%d]=%d\n\n", i, j, f[i-1][j]+f[i][j-1])
			}
		}
	}
	return f[m-1][n-1]
}

func uniquePathsRecursive(m int, n int) int {
	if m-1 == 0 || n-1 == 0 {
		return 1
	} else {
		return uniquePathsRecursive(m-1, n) + uniquePathsRecursive(m, n-1)
	}
}
