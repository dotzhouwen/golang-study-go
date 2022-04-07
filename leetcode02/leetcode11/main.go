package main

import "fmt"

/**
[0:[1], 1:[0, 2], 2:[1,3],3:[2,4],4:[3]]

无向图中连通分量的数目
*/
func countComponents(n int, edges [][]int) int {
	adjTable := make([][]int, n)

	for _, edge := range edges {
		adjTable[edge[0]] = append(adjTable[edge[0]], edge[1])
		adjTable[edge[1]] = append(adjTable[edge[1]], edge[0])
	}

	count := 0
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			dfs(i, &adjTable, &visited)
		}
	}

	return count
}

func dfs(curIndex int, adjTable *[][]int, visited *[]bool) {
	(*visited)[curIndex] = true
	if (*adjTable)[curIndex] != nil {
		for _, i := range (*adjTable)[curIndex] {
			if (*visited)[i] == false {
				dfs(i, adjTable, visited)
			}
		}
	}
}

func main() {
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
	components := countComponents(n, edges)
	fmt.Println(components)
}
