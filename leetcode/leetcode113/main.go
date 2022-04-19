package main

import "fmt"

/**
回溯法 1079. 活字印刷
*/
func numTilePossibilities(tiles string) int {
	chars := []rune(tiles)
	visited := make([]bool, len(chars))
	path := make([]rune, 0)
	set := make(map[string]struct{})

	backtracking(&chars, &visited, &path, set)
	return len(set)
}

func backtracking(chars *[]rune, visited *[]bool, path *[]rune, set map[string]struct{}) {
	if len(*path) > 0 {
		s := string(*path)
		set[s] = struct{}{}
	}

	for i := 0; i < len(*chars); i++ {
		if !(*visited)[i] {
			*path = append(*path, (*chars)[i])
			(*visited)[i] = true
			backtracking(chars, visited, path, set)
			*path = (*path)[:len(*path)-1]
			(*visited)[i] = false
		}
	}
}

func main() {
	/**
	"A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"。
	*/
	title := "AAB"
	possibilities := numTilePossibilities(title)
	fmt.Println(possibilities)
}
