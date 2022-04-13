package main

import (
	"fmt"
	"math"
)

/**
219. 存在重复元素 II
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int][]int, 0)
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		hash[v] = append(hash[v], i)

		if ixs, exists := hash[v]; exists {
			for _, ix := range ixs {
				if i != ix {
					if int(math.Abs(float64(i-ix))) <= k {
						return true
					}
				}
			}
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	duplicate := containsNearbyDuplicate(nums, k)
	fmt.Println(duplicate)
}
