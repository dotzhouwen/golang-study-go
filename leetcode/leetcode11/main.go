package main

import "fmt"

/**
反转字符串中的元音字母

给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。

元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。

输入：s = "hello"
输出："holle"
*/

var hm = map[int32]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func reverseVowels3(s string) string {
	chars := []rune(s)
	left := 0
	right := len(chars) - 1

	for left < right {
		if !IsVowels(chars[left]) {
			left++
			continue
		}
		if !IsVowels(chars[right]) {
			right--
			continue
		}
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}
	return string(chars)
}
func reverseVowels2(s string) string {
	result := []rune(s)

	left := 0
	right := len(s) - 1
	for left < right {
		if !hm[result[left]] {
			left++
			continue
		}
		if !hm[result[right]] {
			right--
			continue
		}
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
	return string(result)
}

func reverseVowels(s string) string {
	chars := []rune(s)
	i := 0
	j := len(chars) - 1

	iDone := false
	jDone := false

	for i < j {
		if !iDone && IsVowels(chars[i]) {
			iDone = true
		} else {
			if !iDone {
				i++
				continue
			}
		}
		if !jDone && IsVowels(chars[j]) {
			jDone = true
		} else {
			if !jDone {
				j--
				continue
			}
		}

		if iDone && jDone {
			chars[i], chars[j] = chars[j], chars[i]
			iDone = false
			jDone = false
			i++
			j--
		}
	}
	return string(chars)
}

func IsVowels(ch int32) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	}
	if ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
		return true
	}
	return false
}

func main() {
	s := "accceiobbbbbbbibbb"
	ret := reverseVowels3(s)
	fmt.Println("ret:", ret)
}
