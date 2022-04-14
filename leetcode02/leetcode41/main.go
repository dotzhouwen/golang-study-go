package main

import (
	"fmt"
	"math"
)

/**
最长无重复子串
*/
func lengthOfLongestSubstring(s string) int {
	chars := []rune(s)
	arr := make([]int, 256)
	left := 0

	maxLen := 0
	for right := 0; right < len(chars); right++ {
		arr[chars[right]-' ']++
		for !check(arr) {
			arr[chars[left]-' ']--
			left++
		}
		length := right - left + 1
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}

func check(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 2 {
			return false
		}
	}
	return true
}

func lengthOfLongestSubstring2(s string) int {
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

func main() {
	s := " "
	substring := lengthOfLongestSubstring(s)
	fmt.Println(substring)
	//
	fmt.Println(' ' - 'a')
}
