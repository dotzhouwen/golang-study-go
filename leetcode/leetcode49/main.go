package main

import "fmt"

func eventualSafeNodes(graph [][]int) []int {

	return nil
}

func main() {
	graph := [][]int{
		{1, 2},
		{2, 3},
		{5},
		{0},
		{5},
		{},
		{},
	}
	fmt.Println(graph)
	nodes := eventualSafeNodes(graph)
	fmt.Println(nodes)
}
