package main

import "fmt"

/**
找到字符串中所有字母异位词
*/
func findAnagrams2(s string, p string) []int {
	sChars := []rune(s)
	pChars := []rune(p)
	ans := make([]int, 0)

	left := 0
	right := 0

	need := make([]int, 26)
	for _, v := range pChars {
		need[v-'a']++
	}

	window := make([]int, 26)
	for right < len(sChars) {
		window[sChars[right]-'a']++
		if right-left+1 == len(pChars) {
			if isEquals(window, need) {
				ans = append(ans, left)
			}
			window[sChars[left]-'a']--
			left++
		}
		right++
	}
	return ans
}

func isEquals(a []int, b []int) bool {
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

func findAnagrams(s string, p string) []int {
	sChars := []rune(s)
	pChars := []rune(p)
	ans := make([]int, 0)

	needHash := make(map[rune]int, 0)
	for _, v := range pChars {
		needHash[v]++
	}

	left := 0
	right := len(p)
	for ; right <= len(sChars); right++ {
		if check(sChars, left, right, &needHash) {
			ans = append(ans, left)
		}
		left++
	}
	return ans
}

func check(sChars []rune, left int, right int, needHash *map[rune]int) bool {
	hash := make(map[rune]int, 0)
	for i := left; i < right; i++ {
		hash[sChars[i]]++
	}
	for k, v := range *needHash {
		if hash[k] != v {
			return false
		}
	}
	return true
}

/**
"cbaebabacd"
"abc"
*/

func main() {
	s := "cbaebabacd"
	p := "abc"
	anagrams := findAnagrams2(s, p)
	fmt.Println(anagrams)
}
