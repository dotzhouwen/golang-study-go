package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	visitedMap := make(map[int]struct{})
	result := make([]int, 0)

	for i := 0; i < len(nums1); i++ {
		v1 := nums1[i]
		flag := false
		for j := 0; j < len(nums2); j++ {
			if nums2[j] == v1 {
				if _, exists := visitedMap[j]; exists {
					continue
				} else {
					flag = true
					visitedMap[j] = struct{}{}
					break
				}
			}
		}
		if flag {
			result = append(result, v1)
		}
	}
	return result
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	ret := intersect(nums1, nums2)
	fmt.Println(ret)
}
