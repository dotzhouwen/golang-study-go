package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
复原ip地址
*/
func restoreIpAddresses(s string) []string {
	var path []string
	var res []string
	backtracking(s, 0, 0, &path, &res)
	return res
}

func backtracking(s string, startIndex int, dotCount int, path *[]string, res *[]string) {
	if startIndex >= len(s) || dotCount >= 3 {
		p := make([]string, len(*path))
		copy(p, *path)

		if startIndex < len(s) {
			s1 := s[startIndex:]
			if isNormalIp(s1) {
				p = append(p, s1)
			} else {
				return
			}
		}

		if len(p) != 4 {
			return
		}

		normalIp := strings.Join(p, ".")
		*res = append(*res, normalIp)
		return
	}
	for i := startIndex; i < len(s); i++ {
		p := s[startIndex : i+1]
		if !isNormalIp(p) {
			continue
		}
		*path = append(*path, p)
		backtracking(s, i+1, dotCount+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

func isNormalIp(s string) bool {
	if strings.HasPrefix(s, "0") {
		if len(s) > 1 {
			return false
		}
		return true
	}
	val, _ := strconv.Atoi(s)
	if val > 255 {
		return false
	}
	return true
}

func main() {
	s := "101023"
	addresses := restoreIpAddresses(s)
	fmt.Println(addresses)
}
