package main

import "fmt"

/**
至多包含 K 个不同字符的最长子串
*/
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	chars := []rune(s)
	hash := make(map[rune]int, 0)

	maxLength := 0
	left := 0
	for right := 0; right < len(chars); right++ {
		hash[chars[right]]++
		for len(hash) > k {
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
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

func main() {
	s := "aa"
	k := 1
	distinct := lengthOfLongestSubstringKDistinct(s, k)
	fmt.Println(distinct)
}
