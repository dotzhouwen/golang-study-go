package main

import (
	"awesomeProject/leetcode/util"
	"fmt"
)

/**
拓扑排序 课程表问题
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses)
	result := make([]int, 0)

	for _, edge := range prerequisites {
		inDegree[edge[0]]++
	}

	queue := util.NewQueue()
	for i, degree := range inDegree {
		if degree == 0 {
			queue.Add(i)
		}
	}

	for !queue.IsEmpty() {
		i := queue.Remove().(int)
		result = append(result, i)

		for _, edge := range prerequisites {
			if edge[1] == i {
				inDegree[edge[0]]--

				if inDegree[edge[0]] == 0 {
					queue.Add(edge[0])
				}
			}
		}
	}
	if len(result) != numCourses {
		return nil
	}
	return result
}

func main() {
	numCourses := 4
	prerequisites := [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2},
	}

	order := findOrder(numCourses, prerequisites)
	fmt.Println(order)
}
