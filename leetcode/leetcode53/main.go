package main

import "fmt"

func minimumSemesters(n int, relations [][]int) int {
	graph := make([][]int, n+1)
	inDegree := make([]int, n+1)

	for _, relation := range relations {
		first, second := relation[0], relation[1]
		graph[first] = append(graph[first], second)
		inDegree[second]++
	}

	queue := make([]int, 0)
	for i, degree := range inDegree {
		if i > 0 && degree == 0 {
			queue = append(queue, i)
		}
	}

	visit := 0
	count := 0
	for len(queue) != 0 {
		size := len(queue)
		count++

		for i := 0; i < size; i++ {
			e := queue[i]
			visit++

			for _, adj := range graph[e] {
				inDegree[adj]--
				if inDegree[adj] == 0 {
					queue = append(queue, adj)
				}
			}
		}
		queue = queue[size:]
	}

	if visit != n {
		return -1
	}
	return count
}

func main() {
	N := 3
	relations := [][]int{{1, 3}, {2, 3}}
	semesters := minimumSemesters(N, relations)

	fmt.Println(semesters)
}
