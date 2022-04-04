package main

import (
	"awesomeProject/leetcode/util"
	"fmt"
)

/**
[][]int{{1, 3}, {2, 3}, {3, 1}}
*/

func minimumSemesters2(n int, relations [][]int) int {
	graph := make([][]int, n+1)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}

	indegree := make([]int, n+1)
	for i := 0; i < len(relations); i++ {
		first, second := relations[i][0], relations[i][1]
		graph[first] = append(graph[first], second)
		indegree[second]++
	}

	var queue []int
	for i := 1; i < n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	ans := 0
	visit := 0

	for len(queue) != 0 {
		ans++
		size := len(queue)

		for i := 0; i < size; i++ {
			cur := queue[i]

			for _, nxt := range graph[cur] {
				indegree[nxt]--
				if 0 == indegree[nxt] {
					queue = append(queue, nxt)
				}
			}
			visit++
		}
		queue = queue[size:]
	}

	if visit != n {
		return -1
	}
	return -1
}

func minimumSemesters(n int, relations [][]int) int {
	inDegree := make([]int, n+1)
	ans := make([]int, 0)

	for _, relation := range relations {
		inDegree[relation[1]]++
	}

	queue := util.NewQueue()
	for i, degree := range inDegree {
		if i > 0 && degree == 0 {
			queue.Add(i)
		}
	}

	count := 0
	for !queue.IsEmpty() {
		queueLen := queue.Len()

		for i := 0; i < queueLen; i++ {
			e := queue.Remove().(int)
			ans = append(ans, e)

			for _, relation := range relations {
				if e == relation[0] {
					inDegree[relation[1]]--

					if inDegree[relation[1]] == 0 {
						queue.Add(relation[1])
					}
				}
			}
		}
		count++
	}
	if len(ans) != n {
		return -1
	}
	return count
}

func main() {
	N := 3
	relations := [][]int{{1, 3}, {2, 3}, {3, 1}}

	semesters := minimumSemesters(N, relations)

	fmt.Println(semesters)
}
