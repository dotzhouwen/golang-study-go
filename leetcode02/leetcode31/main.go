package main

import "fmt"

/**
两个数组的交集
*/
func intersection(nums1 []int, nums2 []int) []int {
	hash1 := make(map[int]struct{}, 0)
	for _, v := range nums1 {
		hash1[v] = struct{}{}
	}

	hash2 := make(map[int]struct{}, 0)
	for _, v := range nums2 {
		hash2[v] = struct{}{}
	}

	var res []int
	for k, _ := range hash1 {
		if _, exists := hash2[k]; exists {
			res = append(res, k)
		}
	}

	return res
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	ints := intersection(nums1, nums2)
	fmt.Println(ints)
}
