package main

import "fmt"

/**
滑动窗口 ，最小覆盖子串
*/
func minWindow(s string, t string) string {
	sChars := []rune(s)
	tChars := []rune(t)

	charMap := make(map[rune]int, 128)
	for _, ch := range tChars {
		if _, exists := charMap[ch]; exists {
			charMap[ch] = charMap[ch] + 1
		} else {
			charMap[ch] = 1
		}
	}

	left := 0
	right := 0

	minLen := len(sChars)
	minLenStr := s

	for left <= right && right < len(sChars) {
		ch := sChars[right]
		if _, exists := charMap[ch]; exists {
			charMap[ch]--
		}
		for check(charMap) {
			length := right - left + 1
			if length < minLen {
				minLen = length
				minLenStr = string(sChars[left : right+1])
			}

			leftChar := sChars[left]
			if _, exists := charMap[leftChar]; exists {
				charMap[leftChar]++
			}
			left++
		}
		right++
	}

	if left == 0 && right == len(s) {
		return ""
	}
	return minLenStr
}

func check(charMap map[rune]int) bool {
	for _, v := range charMap {
		if v > 0 {
			return false
		}
	}
	return true
}

func contains(s []rune, t []rune) bool {
	visited := make([]bool, len(s))
	for _, tch := range t {
		exists := false
		for i, sch := range s {
			if sch == tch && !visited[i] {
				visited[i] = true
				exists = true
				break
			}
		}
		if !exists {
			return false
		}
	}
	return true
}
func main() {
	s := "a"
	t := "aa"
	window := minWindow(s, t)
	fmt.Println(window)
}
