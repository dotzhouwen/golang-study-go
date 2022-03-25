package main

import (
	"fmt"
	"strings"
)

/**
验证回文串

给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

输入: "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串
*/

func isPalindrome(s string) bool {
	if s == "" || len(s) == 1 {
		return true
	}

	s = strings.ToLower(s)
	chars := []rune(s)
	i := 0
	j := len(chars) - 1

	var flag = true
	for i < j {
		for i < j && !ischar(chars[i]) {
			i++
		}
		for i < j && !ischar(chars[j]) {
			j--
		}

		if i > j {
			flag = false
		}

		fmt.Printf("i:%d, j:%d, ch[i]=%d, ch[j]=%d\n", i, j, chars[i], chars[j])
		if chars[i] != chars[j] {
			flag = false
			break
		}
		i++
		j--
	}

	return flag
}

func ischar(ch int32) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
		return true
	}
	return false
}

func main() {
	s := "A man, a plan, a canal: Panama"
	ret := isPalindrome(s)
	fmt.Println(ret)
}
