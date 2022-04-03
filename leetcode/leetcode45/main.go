package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {
	graph := make(map[int][]int, 0)
	visited := make(map[int]struct{})

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]

		if _, existU := graph[u]; existU {
			if _, existV := graph[v]; existV {
				visited = make(map[int]struct{})
				if dfs(graph, u, v, visited) {
					return edge
				}
			}
		}

		if _, exists := graph[u]; exists {
			graph[u] = append(graph[u], v)
		} else {
			adj := make([]int, 0)
			adj = append(adj, v)
			graph[u] = adj
		}

		if _, exists := graph[v]; exists {
			graph[v] = append(graph[v], u)
		} else {
			adj := make([]int, 0)
			adj = append(adj, u)
			graph[v] = adj
		}
	}

	return nil
}

func dfs(graph map[int][]int, source int, target int, visited map[int]struct{}) bool {
	if source == target {
		return true
	}

	visited[source] = struct{}{}
	if adjEdges, exists := graph[source]; exists {
		for _, adj := range adjEdges {
			if _, ex := visited[adj]; !ex {
				if dfs(graph, adj, target, visited) {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	redundant := findRedundantConnection(edges)

	fmt.Println(redundant)
}
