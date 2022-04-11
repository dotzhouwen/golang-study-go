package main

import "fmt"

/**
159. 至多包含两个不同字符的最长子串
*/
func lengthOfLongestSubstringTwoDistinct(s string) int {
	chars := []rune(s)
	hash := make(map[rune]int)

	left := 0
	maxLen := 0
	for right := 0; right < len(chars); right++ {
		hash[chars[right]]++
		for len(hash) > 2 {
			if cnt, exists := hash[chars[left]]; exists {
				if cnt == 1 {
					delete(hash, chars[left])
				} else {
					hash[chars[left]]--
				}
			}
			left++
		}
		length := right - left + 1
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}

func main() {
	s := "ccaabbb"
	distinct := lengthOfLongestSubstringTwoDistinct(s)
	fmt.Println(distinct)
}
