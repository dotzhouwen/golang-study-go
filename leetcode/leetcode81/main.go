package main

import "fmt"

/**
树的直径
*/
func treeDiameter(edges [][]int) int {
	adjTable := make([][]int, len(edges)+1)

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		adjTable[u] = append(adjTable[u], v)
		adjTable[v] = append(adjTable[v], u)
	}

	last, _ := bfs(adjTable, 0, len(edges)+1)
	_, diameter := bfs(adjTable, last, len(edges)+1)

	return diameter
}

func bfs(adjTable [][]int, first int, length int) (int, int) {
	queue := make([]int, 0)
	queue = append(queue, first)

	var last int
	diameter := -1

	visited := make([]bool, length)
	for len(queue) > 0 {
		size := len(queue)
		diameter++
		for i := 0; i < size; i++ {
			e := queue[0]
			queue = queue[1:]
			visited[e] = true
			last = e

			for _, v := range adjTable[e] {
				if !visited[v] {
					queue = append(queue, v)
				}
			}
		}
	}
	return last, diameter
}

func main() {
	edges := [][]int{{0, 1}, {0, 2}}
	diameter := treeDiameter(edges)
	fmt.Println(diameter)
}
