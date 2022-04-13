package main

import "fmt"

/*
两数之和
*/
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, 0)
	for i, v := range nums {
		if ix, exists := hash[target-v]; exists {
			return []int{ix, i}
		} else {
			hash[v] = i
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 100
	sum := twoSum(nums, target)
	fmt.Println(sum)
}
