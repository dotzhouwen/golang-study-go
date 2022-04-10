package main

import "fmt"

func shortestBridge(grid [][]int) int {
	return 0
}

func main() {
	A := [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}
	bridge := shortestBridge(A)
	fmt.Println(bridge)
}
