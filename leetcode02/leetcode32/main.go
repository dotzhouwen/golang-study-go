package main

import "fmt"

/**
快乐数 leetcode 202
*/
func isHappy(n int) bool {
	hash := make(map[int]struct{})
	next := n
	for {
		res := 0
		hash[next] = struct{}{}
		for next > 0 {
			res += (next % 10) * (next % 10)
			next = next / 10
		}
		if res == 1 {
			return true
		}
		if _, exists := hash[res]; exists {
			return false
		}
		next = res
	}
}

func main() {
	n := 1
	happy := isHappy(n)
	fmt.Println(happy)
}
