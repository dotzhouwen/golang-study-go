package main

import "fmt"

/**
字符串的排列
*/
func checkInclusion(s1 string, s2 string) bool {
	chars1 := []rune(s1)
	chars2 := []rune(s2)

	need := make([]rune, 26)
	for _, v := range chars1 {
		need[v-'a']++
	}

	left := 0
	right := 0

	window := make([]rune, 26)
	for ; right < len(chars2); right++ {
		window[chars2[right]-'a']++
		if right-left+1 == len(s1) {
			if isEquals(window, need) {
				return true
			}
			window[chars2[left]-'a']--
			left++
		}
	}
	return false
}

func isEquals(a []rune, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	s1 := "ab"
	s2 := "eidboaoo"
	inclusion := checkInclusion(s1, s2)
	fmt.Println(inclusion)
}
