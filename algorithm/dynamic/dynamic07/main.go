package main

import (
	"fmt"
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
			if j == 0 {
				minValue := arrayMin(f[i-1][1:])
				f[i][j] = minValue + grid[i][j]

			} else if j == len(f[i])-1 {
				minValue := arrayMin(f[i-1][:len(f[i])-1])
				f[i][j] = minValue + grid[i][j]

			} else {
				var tmp []int
				tmp = append(tmp, f[i-1][:j]...)
				tmp = append(tmp, f[i-1][j+1:len(f[i])]...)

				minValue := arrayMin(tmp)
				f[i][j] = minValue + grid[i][j]
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
