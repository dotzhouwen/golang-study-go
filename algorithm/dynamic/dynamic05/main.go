package main

import (
	"fmt"
	"math"
)

/**
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1

输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为11（即，2+3+5+1= 11）。


*/
func minimumTotal(triangle [][]int) int {
	f := make([][]int, len(triangle))
	for i := range f {
		f[i] = make([]int, len(triangle[i]))
	}

	f[0][0] = triangle[0][0]
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f[i]); j++ {
			if i == 0 && j == 0 {
				f[i][j] = triangle[i][i]
			} else if j == 0 {
				f[i][j] = f[i-1][j] + triangle[i][j]
			} else {
				lastLineLen := len(f[i-1])
				if j <= lastLineLen-1 {
					f[i][j] = int(math.Min(float64(f[i-1][j-1]+triangle[i][j]), float64(f[i-1][j]+triangle[i][j])))
				} else {
					f[i][j] = f[i-1][j-1] + triangle[i][j]
				}
			}
		}
	}

	lastLine := f[len(f)-1]
	min := lastLine[0]
	for i := 0; i < len(lastLine); i++ {
		if min > lastLine[i] {
			min = lastLine[i]
		}
	}
	return min
}

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	fmt.Println(triangle)

	ret := minimumTotal(triangle)
	fmt.Println(ret)
}
