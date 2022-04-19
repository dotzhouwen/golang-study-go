package main

import (
	"fmt"
	"sort"
)

/**
698. 划分为k个相等的子集
*/
func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	target := sum / k
	if target*k != sum {
		return false
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	if nums[0] > target {
		return false
	}

	newNums := make([]int, 0)
	for _, num := range nums {
		if num != target {
			newNums = append(newNums, num)
		}
	}
	sub := len(nums) - len(newNums)

	bucket := make([]int, k-sub)
	return backtracking(&newNums, 0, k-sub, target, &bucket)
}

func backtracking(nums *[]int, cur int, k int, target int, bucket *[]int) bool {
	if cur == len(*nums) {
		return true
	}

	for i := 0; i < k; i++ {
		v := (*nums)[cur]
		if (*bucket)[i]+v <= target {
			(*bucket)[i] += v
			if backtracking(nums, cur+1, k, target, bucket) {
				return true
			}
			(*bucket)[i] -= v
		}
	}
	return false
}

func main() {
	nums := []int{10, 1, 10, 9, 6, 1, 9, 5, 9, 10, 7, 8, 5, 2, 10, 8}
	k := 11
	subsets := canPartitionKSubsets(nums, k)
	fmt.Println(subsets)
}
