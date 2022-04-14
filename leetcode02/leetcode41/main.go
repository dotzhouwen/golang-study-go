package main

import "fmt"

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

func main() {
	s := " "
	substring := lengthOfLongestSubstring(s)
	fmt.Println(substring)
	//
	fmt.Println(' ' - 'a')
}
