package main

import "fmt"

/**
分割回文串
*/
func partition(s string) [][]string {
	path := make([]string, 0)
	res := make([][]string, 0)
	backtracking(s, 0, &path, &res)

	return res
}

func backtracking(s string, startIndex int, path *[]string, res *[][]string) {
	if startIndex >= len(s) {
		p := make([]string, len(*path))
		copy(p, *path)
		*res = append(*res, p)
		return
	}
	for i := startIndex; i < len(s); i++ {
		if !isPalindrome(s[startIndex : i+1]) {
			continue
		}
		*path = append(*path, s[startIndex:i+1])
		backtracking(s, i+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func combine(n int, k int) [][]int {
	var path []int
	var res [][]int

	combingBacktracking(n, k, 1, &path, &res)
	return res
}

func combingBacktracking(n int, k int, startIndex int, path *[]int, res *[][]int) {
	if len(*path) == k {
		p := make([]int, len(*path))
		copy(p, *path)
		*res = append(*res, p)
		return
	}

	for i := startIndex; i <= n; i++ {
		*path = append(*path, i)
		combingBacktracking(n, k, i+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	s := "a"
	results := partition(s)
	for _, result := range results {
		fmt.Println(result)
	}
}
