package main

import (
	"fmt"
	"unicode"
)

/**
字母大小写全排列
*/
func letterCasePermutation(s string) []string {
	var path []rune
	var res []string

	dfs([]rune(s), 0, &path, &res)
	return res
}

func dfs(chars []rune, startIndex int, path *[]rune, res *[]string) {
	if len(*path) >= len(chars) {
		s := string(*path)
		*res = append(*res, s)
		return
	}

	for i := startIndex; i < len(chars); i++ {
		if !unicode.IsDigit(chars[i]) {
			lower := unicode.ToLower(chars[i])
			*path = append(*path, lower)
			dfs(chars, i+1, path, res)
			*path = (*path)[:len(*path)-1]

			upper := unicode.ToUpper(chars[i])
			*path = append(*path, upper)
			dfs(chars, i+1, path, res)
			*path = (*path)[:len(*path)-1]
		} else {
			*path = append(*path, chars[i])
			dfs(chars, i+1, path, res)
			*path = (*path)[:len(*path)-1]
		}
	}
}

func main() {
	s := "abcd"
	permutation := letterCasePermutation(s)
	fmt.Println(permutation)
}
