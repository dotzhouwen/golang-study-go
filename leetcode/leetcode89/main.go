package main

import (
	"fmt"
	"math"
)

func findLengthOfLCIS(nums []int) int {
	length := len(nums)
	i := 0
	j := 0

	res := 0
	for j < length {
		if j > 0 && nums[j] <= nums[j-1] {
			i = j
		}
		j++
		res = int(math.Max(float64(res), float64(j-i)))
	}
	return res
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	lcis := findLengthOfLCIS(nums)
	fmt.Println(lcis)
}
