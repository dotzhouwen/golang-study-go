package main

import "fmt"

func letterCombinations(digits string) []string {
	letterMap := map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	if digits == "" {
		return nil
	}

	digitChar := []rune(digits)
	path := make([]rune, 0)
	res := make([]string, 0)

	backtracking(0, digitChar, letterMap, &path, &res)
	return res
}

func backtracking(startIndex int, digits []rune, letterMap map[rune][]rune, path *[]rune, res *[]string) {
	if len(*path) == len(digits) {
		p := make([]rune, len(digits))
		copy(p, *path)
		*res = append(*res, string(p))
	}

	if startIndex < len(digits) {
		if letters, ok := letterMap[digits[startIndex]]; ok {
			for _, letter := range letters {
				*path = append(*path, letter)
				backtracking(startIndex+1, digits, letterMap, path, res)
				*path = (*path)[:len(*path)-1]
			}
		}
	}
}

func main() {
	digits := "2345"
	combinations := letterCombinations(digits)
	fmt.Println(combinations)
}
