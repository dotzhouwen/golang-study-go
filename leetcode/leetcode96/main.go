package main

import (
	"fmt"
	"math"
)

/**
无重复的最长子字符串
*/
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	chars := []rune(s)
	charMap := make(map[rune]int, 128)

	j := 0
	ans := 0
	for i := 0; i < len(chars); i++ {
		v := chars[i]
		charMap[v]++
		for charMap[v] > 1 {
			charMap[chars[j]]--
			j++
		}
		ans = int(math.Max(float64(ans), float64(i-j+1)))
	}
	return ans
}

func check(hash map[rune]int) bool {
	for _, ch := range hash {
		if ch > 1 {
			return false
		}
	}
	return true
}

func main() {
	s := "abcda"
	substring := lengthOfLongestSubstring(s)
	fmt.Println(substring)
}
