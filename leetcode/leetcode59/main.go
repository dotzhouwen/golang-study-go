package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	var track []int
	var result [][]int
	backTree(n, k, 1, &track, &result)
	return result
}

func backTree(n int, k int, startIndex int, track *[]int, result *[][]int) {
	if len(*track) == k {
		var sum int
		tmp := make([]int, k)
		for k, v := range *track {
			sum += v
			tmp[k] = v
		}
		if sum == n {
			*result = append(*result, tmp)
		}
		return
	}

	for i := startIndex; i <= 9-(k-len(*track))+1; i++ {
		*track = append(*track, i)
		backTree(n, k, i+1, track, result)
		*track = (*track)[:len(*track)-1]
	}
}

func main() {
	k := 3
	n := 9
	sum3 := combinationSum3(k, n)
	fmt.Println(sum3)
}
