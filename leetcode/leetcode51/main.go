package main

import (
	"awesomeProject/leetcode/util"
	"fmt"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
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
		e := queue.Remove().(int)
		result = append(result, e)

		for _, edge := range prerequisites {
			if e == edge[1] {
				inDegree[edge[0]]--

				if inDegree[edge[0]] == 0 {
					queue.Add(edge[0])
				}
			}
		}
	}

	if len(result) != numCourses {
		return false
	}
	return true
}

func main() {
	numCourses := 4
	prerequisites := [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2},
	}
	finish := canFinish(numCourses, prerequisites)
	fmt.Println(finish)
}
