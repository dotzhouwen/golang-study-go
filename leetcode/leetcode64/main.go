package main

import "fmt"

func partition(s string) [][]string {
	var tmpString []string
	var res [][]string

	backtracking(s, tmpString, 0, &res)
	return res
}

func backtracking(s string, tmpString []string, startIndex int, res *[][]string) {
	if startIndex == len(s) {
		t := make([]string, len(tmpString))
		copy(t, tmpString)
		*res = append(*res, t)
	}

	for i := startIndex; i < len(s); i++ {
		if isPartition(s, startIndex, i) {
			tmpString = append(tmpString, s[startIndex:i+1])
		} else {
			continue
		}
		backtracking(s, tmpString, i+1, res)
		tmpString = tmpString[:len(tmpString)-1]
	}
}

func isPartition(s string, startIndex int, end int) bool {
	left := startIndex
	right := end
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "abcba"
	i := partition(s)
	fmt.Println(i)
}
