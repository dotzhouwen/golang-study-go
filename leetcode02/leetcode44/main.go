package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	hash := make(map[int]int, 0)
	for _, v := range nums {
		hash[v]++
	}

	values := make([]int, 0)
	for _, v := range hash {
		values = append(values, v)
	}

	res := make([]int, 0)
	sort.Ints(values)

	visited := make(map[int]bool)
	for i := len(values) - 1; i >= len(values)-k; i-- {
		for k, v := range hash {
			if v == values[i] && !visited[k] {
				res = append(res, k)
				visited[k] = true
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 2}
	k := 2
	frequent := topKFrequent(nums, k)
	fmt.Println(frequent)
}
