package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	child := 0
	sIndex := 0

	for child < len(g) && sIndex < len(s) {
		if s[sIndex] >= g[child] {
			child++
		}
		sIndex++
	}

	return child
}

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 3, 3}

	children := findContentChildren(g, s)
	fmt.Println(children)
}
