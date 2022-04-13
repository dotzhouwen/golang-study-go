package main

import "fmt"

/**
205. 同构字符串
*/
func isIsomorphic(s string, t string) bool {
	sChars := []rune(s)
	tChars := []rune(t)

	if len(sChars) != len(tChars) {
		return false
	}

	hash1 := make(map[rune]rune, 0)
	hash2 := make(map[rune]rune, 0)

	for i := 0; i < len(sChars); i++ {
		ch := sChars[i]
		if v, exists := hash1[ch]; exists {
			if v != tChars[i] {
				return false
			}
		} else {
			hash1[ch] = tChars[i]
		}
	}

	for i := 0; i < len(tChars); i++ {
		ch := tChars[i]
		if v, exists := hash2[ch]; exists {
			if v != sChars[i] {
				return false
			}
		} else {
			hash2[ch] = sChars[i]
		}
	}
	return true
}

func main() {
	s := "paper"
	t := "title"
	isomorphic := isIsomorphic(s, t)
	fmt.Println(isomorphic)
}
