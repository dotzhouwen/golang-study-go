package main

import "fmt"

/**
平行课程
*/
func minimumSemesters(n int, relations [][]int) int {
	inDegree := make([]int, n+1)
	for _, relation := range relations {
		inDegree[relation[1]]++
	}

	queue := make([]int, 0)
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	total := 0
	for len(queue) > 0 {
		size := len(queue)
		count++

		for i := 0; i < size; i++ {
			e := queue[0]
			queue = queue[1:]
			total++

			for _, relation := range relations {
				if e == relation[0] {
					inDegree[relation[1]]--
					if inDegree[relation[1]] == 0 {
						queue = append(queue, relation[1])
					}
				}
			}
		}
	}

	if total != n {
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
