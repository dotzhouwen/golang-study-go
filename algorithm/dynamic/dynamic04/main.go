package main

import (
	"fmt"
	"math"
)

/**
给定一个包含非负整数的 mxn网格grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例 1：


输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。

*/

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				f[i][j] = grid[0][0]
			} else if i == 0 {
				f[i][j] = f[i][j-1] + grid[i][j]
			} else if j == 0 {
				f[i][j] = f[i-1][j] + grid[i][j]
			} else {
				f[i][j] = int(math.Min(float64(f[i-1][j]+grid[i][j]), float64(f[i][j-1]+grid[i][j])))
			}
		}
	}
	return f[m-1][n-1]
}

func main() {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	ret := minPathSum(grid)
	fmt.Println(ret)
}
