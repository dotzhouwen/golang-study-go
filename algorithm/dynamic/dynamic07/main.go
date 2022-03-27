package main

import (
	"fmt"
	"math"
)

func minFallingPathSum(grid [][]int) int {
	f := make([][]int, len(grid))
	for i := range f {
		f[i] = make([]int, len(grid[i]))
	}

	for i := 0; i < len(grid[0]); i++ {
		f[0][i] = grid[0][i]
	}

	for i := 1; i < len(f); i++ {
		for j := 0; j < len(f[i]); j++ {
			f[i][j] = math.MaxInt
			val := grid[i][j]
			for p := 0; p < len(f[i]); p++ {
				if j != p {
					f[i][j] = int(math.Min(float64(f[i][j]), float64(f[i-1][p]+val)))
				}
			}
		}
	}

	min := f[len(f)-1][0]
	for i := 0; i < len(f[len(f)-1]); i++ {
		if f[len(f)-1][i] < min {
			min = f[len(f)-1][i]
		}
	}
	return min
}

func arrayMin(arr []int) (min int) {
	min = arr[0]
	for i := range arr {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	sum := minFallingPathSum(arr)
	fmt.Println(sum)
}
