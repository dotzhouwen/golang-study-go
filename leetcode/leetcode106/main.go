package main

import "fmt"

/**
dfs 547. 省份数量
*/
func findCircleNum(isConnected [][]int) int {
	m := len(isConnected)
	visited := make([]bool, m)

	count := 0
	for i := 0; i < m; i++ {
		if !visited[i] {
			visited[i] = true
			dfs(&isConnected, &visited, i)
			count++
		}
	}
	return count
}

func dfs(isConnected *[][]int, visited *[]bool, i int) {
	for j, _ := range (*isConnected)[i] {
		if !(*visited)[j] && (*isConnected)[i][j] == 1 {
			(*visited)[j] = true
			dfs(isConnected, visited, j)
		}
	}
}

func main() {
	/**
	0-3
	1-2
	2-3
	2-4
	3-2
	*/
	isConnected := [][]int{

		{1, 1, 0}, {1, 1, 0}, {0, 0, 1},
	}
	num := findCircleNum(isConnected)
	fmt.Println(num)
}
