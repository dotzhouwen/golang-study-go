package main

import (
	"fmt"
	"math"
)

/**

931. 下降路径最小和

给你一个 n x n 的 方形 整数数组matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。

下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。
具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。


*/

func minFallingPathSum(matrix [][]int) int {
	f := make([][]int, len(matrix))
	for i := range f {
		f[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix[0]); i++ {
		f[0][i] = matrix[0][i]
	}

	for i := 1; i < len(f); i++ {
		for j := 0; j < len(f[i]); j++ {
			if j == 0 && j != len(f[i])-1 {
				f[i][j] = int(math.Min(float64(f[i-1][j]+matrix[i][j]), float64(f[i-1][j+1]+matrix[i][j])))
			} else if j != 0 && j == len(f[i])-1 {
				f[i][j] = int(math.Min(float64(f[i-1][j]+matrix[i][j]), float64(f[i-1][j-1]+matrix[i][j])))
			} else if j == 0 && j == len(f[i])-1 {
				f[i][j] = f[i-1][j] + matrix[i][j]
			} else {
				f[i][j] = int(math.Min(math.Min(float64(f[i-1][j]+matrix[i][j]), float64(f[i-1][j+1]+matrix[i][j])), float64(f[i-1][j-1]+matrix[i][j])))
			}
		}
	}

	lastLine := f[len(f)-1]
	min := lastLine[0]
	for i := 0; i < len(lastLine); i++ {
		if lastLine[i] < min {
			min = lastLine[i]
		}
	}
	return min
}

func arrayMin(arr []int) int {
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func main() {
	matrix := [][]int{
		{-84, -36, 2},
		{87, -79, 10},
		{42, 10, 63},
	}

	ret := minFallingPathSum(matrix)
	fmt.Println(ret)
}
