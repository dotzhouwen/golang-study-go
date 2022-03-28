package main

import (
	"fmt"
	"strings"
	"unicode"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	strchars := make([][]rune, len(strs))
	for i, s := range strs {
		strchars[i] = []rune(s)
	}
	firstChars := strchars[0]
	i := 0
	for ; i < len(firstChars); i++ {
		ch := firstChars[i]
		flag := false
		for j := 1; j < len(strchars); j++ {
			chars := strchars[j]
			if i < len(chars) {
				if chars[i] == ch {
					continue
				} else {
					flag = true
					break
				}
			} else {
				flag = true
				break
			}
		}
		if flag {
			return string(firstChars[:i])
		}
	}
	if i == len(firstChars) {
		return string(firstChars)
	}
	return ""
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	hchars := []rune(haystack)
	nchars := []rune(needle)

	i := 0
	j := 0

	for i < len(hchars) && j < len(nchars) {
		if hchars[i] == nchars[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j == len(nchars) {
		return i - j
	}
	return -1
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	chars := []int32(s)

	result := 0
	negative := false
	for i := 0; i < len(chars); i++ {
		v := chars[i]
		if i == 0 && v == '-' {
			negative = true
			continue
		} else if i == 0 && v == '+' {
			negative = false
			continue
		}

		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
			if result >= 2147483648 {
				if negative {
					return -2147483648
				} else {
					return 2147483647
				}
			}
		} else {
			break
		}
	}
	if negative {
		result = -result
	}
	return result
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	chars := []rune(s)
	left := 0
	right := len(chars) - 1

	for left < right {
		if !(unicode.IsDigit(chars[left]) || unicode.IsLetter(chars[left])) {
			left++
			continue
		}
		if !(unicode.IsDigit(chars[right]) || unicode.IsLetter(chars[right])) {
			right--
			continue
		}
		if chars[left] != chars[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isAnagram2(s string, t string) bool {
	cs := []byte(s)
	ct := []byte(t)
	var count [26]int

	for i := 0; i < len(cs); i++ {
		count[cs[i]-'a']++
	}

	for i := 0; i < len(ct); i++ {
		count[ct[i]-'a']--
	}

	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			return false
		}
	}
	return true
}

func isAnagram(s string, t string) bool {
	chars1 := []byte(s)
	chars2 := []byte(t)
	charMap1 := make(map[byte]int)
	charMap2 := make(map[byte]int)

	for i := 0; i < len(chars1); i++ {
		charMap1[chars1[i]]++
	}
	for i := 0; i < len(chars2); i++ {
		charMap2[chars2[i]]++
	}

	if len(charMap1) != len(charMap2) {
		return false
	}
	for k, v := range charMap1 {
		if v2, exists := charMap2[k]; exists {
			if v2 != v {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func main() {
	strs := []string{"ab", "ab", "ab"}
	prefix := longestCommonPrefix(strs)
	fmt.Println(prefix)
}
