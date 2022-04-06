package main

import (
	"fmt"
)

/**
22. 括号生成
*/
func generateParenthesis(n int) []string {
	var path []rune
	var res []string

	var chars []rune
	for i := 0; i < n; i++ {
		chars = append(chars, '(')
		chars = append(chars, ')')
	}

	used := make([]bool, len(chars))
	backtracking(chars, used, &path, &res)
	return res
}

func backtracking(chars []rune, used []bool, path *[]rune, res *[]string) {
	if len(*path) >= len(chars) {
		if isValidBracket(*path) {
			s := string(*path)
			*res = append(*res, s)
		}
		return
	}
	uset := make(map[rune]struct{})
	for i := 0; i < len(chars); i++ {
		if _, exists := uset[chars[i]]; exists {
			continue
		}
		if !used[i] {
			*path = append(*path, chars[i])
			uset[chars[i]] = struct{}{}
			used[i] = true

			backtracking(chars, used, path, res)
			*path = (*path)[:len(*path)-1]
			used[i] = false
		}
	}
}

func isValidBracket(path []rune) bool {
	stack := make([]rune, 0)
	for _, ch := range path {
		if ch == '(' {
			stack = append(stack, ch)
		} else {
			if len(stack) > 0 {
				stack = stack[1:]
			}
		}
	}
	return len(stack) == 0
}

func main() {
	parenthesis := generateParenthesis(5)
	for _, item := range parenthesis {
		fmt.Println(item)
	}
}
