package main

import "fmt"

/**
1456. 定长子串中元音的最大数目
*/
func maxVowels(s string, k int) int {
	chars := []rune(s)
	vowelCount := 0

	i := 0
	for ; i < k; i++ {
		if isVowel(chars[i]) {
			vowelCount++
		}
	}

	temp := vowelCount
	for ; i < len(chars); i++ {
		if isVowel(chars[i]) {
			temp++
		}
		if isVowel(chars[i-k]) {
			temp--
		}
		if temp > vowelCount {
			vowelCount = temp
		}
	}

	return vowelCount
}

func isVowel(ch rune) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	}
	return false
}

func main() {
	s := "leetcode"
	k := 3
	vowels := maxVowels(s, k)
	fmt.Println(vowels)
}
