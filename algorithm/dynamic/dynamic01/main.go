package main

import (
	"fmt"
	"math"
)

func coinChange(A []int, M int) int {
	f := make([]int, M+1)
	n := len(A)

	f[0] = 0
	for i := 1; i <= M; i++ {
		f[i] = math.MaxInt
		for j := 0; j < n; j++ {
			if i >= A[j] && f[i-A[j]] != math.MaxInt {
				f[i] = int(math.Min(float64(f[i-A[j]]+1), float64(f[i])))
			}
		}
	}

	fmt.Println(f)

	if f[M] == math.MaxInt {
		f[M] = -1
	}
	return f[M]
}

func main() {
	nums := []int{2, 5, 7}
	ret := coinChange(nums, 27)
	fmt.Println("ret:", ret)
}
