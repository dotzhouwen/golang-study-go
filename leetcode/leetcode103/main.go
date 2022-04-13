package main

import (
	"fmt"
	"sort"
)

/**
字母异位词分组
*/
func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)
	for _, str := range strs {
		chars := []rune(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		key := string(chars)
		if _, exists := hash[key]; exists {
			hash[key] = append(hash[key], str)
		} else {
			hash[key] = []string{str}
		}
	}

	ans := make([][]string, 0)
	for _, v := range hash {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	anagrams := groupAnagrams(strs)
	fmt.Println(anagrams)
}
