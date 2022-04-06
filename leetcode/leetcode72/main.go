package main

import "fmt"

/**
22. 括号生成 解法2
*/
func generateParenthesis(n int) []string {
	var path string
	var res []string

	dfs(n, path, &res, 0, 0)

	return res
}

func dfs(n int, path string, res *[]string, open int, close int) {
	if open > n || close > open {
		return
	}
	if len(path) >= 2*n {
		*res = append(*res, path)
		return
	}

	dfs(n, path+"(", res, open+1, close)
	dfs(n, path+")", res, open, close+1)
}

func main() {
	parenthesis := generateParenthesis(3)
	fmt.Println(parenthesis)
}
