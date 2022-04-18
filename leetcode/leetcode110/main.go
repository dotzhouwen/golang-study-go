package main

import (
	"fmt"
	"sort"
)

/**
火柴拼正方形
*/
func makesquare(matchsticks []int) bool {
	sum := 0
	for _, num := range matchsticks {
		sum += num
	}

	sideLength := sum / 4
	if sideLength*4 != sum {
		return false
	}

	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})

	size := make([]int, 4)
	return backtracking(&matchsticks, 0, sideLength, &size)
}

func backtracking(matchsticks *[]int, startIndex int, sideLength int, size *[]int) bool {
	if startIndex == len(*matchsticks) {
		if (*size)[0] == (*size)[1] && (*size)[1] == (*size)[2] && (*size)[2] == (*size)[3] {
			return true
		}
		return false
	}

	for i := 0; i < 4; i++ {
		if (*size)[i]+(*matchsticks)[startIndex] > sideLength {
			continue
		}
		(*size)[i] += (*matchsticks)[startIndex]
		if backtracking(matchsticks, startIndex+1, sideLength, size) {
			return true
		}
		(*size)[i] -= (*matchsticks)[startIndex]
	}
	return false
}

func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 102}
	b := makesquare(nums)
	fmt.Println(b)
}
