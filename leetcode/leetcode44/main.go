package main

import "fmt"

func countComponents(n int, edges [][]int) int {
	adjVex := make([][]int, n)
	for i := 0; i < n; i++ {
		adjVex[i] = []int{}
	}

	for _, edge := range edges {
		adj0 := adjVex[edge[0]]
		adj0 = append(adj0, edge[1])
		adjVex[edge[0]] = adj0

		adj1 := adjVex[edge[1]]
		adj1 = append(adj1, edge[0])
		adjVex[edge[1]] = adj1
	}

	visited := make([]bool, n)
	count := 0

	for v, _ := range adjVex {
		if !visited[v] {
			dfs(adjVex, v, visited)
			count++
		}
	}

	return count
}

func dfs(adjVex [][]int, v int, visited []bool) {
	visited[v] = true
	vertex := adjVex[v]
	for _, n := range vertex {
		if !visited[n] {
			dfs(adjVex, n, visited)
		}
	}
}

func main() {
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
	components := countComponents(n, edges)

	fmt.Println(components)
}
