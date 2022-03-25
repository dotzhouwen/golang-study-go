package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	curLen := m
	for i := 0; i < n; i++ {
		v := nums2[i]
		j := 0
		for j < curLen && v > nums1[j] {
			j++
		}
		if j == curLen {
			nums1[curLen] = v
		} else {
			for k := curLen - 1; k >= j; k-- {
				nums1[k+1] = nums1[k]
			}
			nums1[j] = v
		}
		curLen++
	}
}

func main() {
	nums1 := []int{5, 6, 7, 9, 10, 0, 0, 0, 0}
	m := 5
	nums2 := []int{2, 5, 6, 11}
	n := 4

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
