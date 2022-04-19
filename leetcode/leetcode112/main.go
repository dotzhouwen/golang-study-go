package main

import (
	"fmt"
	"strings"
)

/**
练习：拆分字符串使唯一子字符串的数目最大
*/
func maxUniqueSplit(s string) int {
	chars := []rune(s)
	maxLen := new(int)
	hash := make(map[string]struct{}, 0)
	backtracking(&chars, 0, hash, maxLen)
	return *maxLen
}

func backtracking(chars *[]rune, startIndex int, hash map[string]struct{}, maxLen *int) {
	if startIndex == len(*chars) {
		if len(hash) > *maxLen {
			*maxLen = len(hash)
		}
		return
	}
	for i := startIndex; i < len(*chars); i++ {
		chs := (*chars)[startIndex : i+1]
		s := string(chs)
		if _, exists := hash[s]; !exists {
			hash[s] = struct{}{}
			backtracking(chars, i+1, hash, maxLen)
			delete(hash, s)
		}
	}
}

func printHashKeys(hash map[string]struct{}) {
	builder := strings.Builder{}
	for k, _ := range hash {
		builder.WriteString(fmt.Sprintf("%s ", k))
	}
	fmt.Println(builder.String())
}

func main() {
	s := "ababccc"
	split := maxUniqueSplit(s)
	fmt.Println(split)
}
