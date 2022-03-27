package main

import (
	stack2 "awesomeProject/algorithm/dfs/dfs01/stack"
	"fmt"
)

type void struct{}

var member void

func main() {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "C", "D"},
		"C": {"A", "B", "D", "E"},
		"D": {"B", "C", "E", "F"},
		"E": {"C", "D"},
		"F": {"D"},
	}

	DFS(graph, "A")
}

func DFS(graph map[string][]string, start string) {
	stack := stack2.NewStack()
	stack.Push(start)

	closed := make(map[string]void)
	closed[start] = member

	for !stack.Empty() {
		vertex := stack.Pop().(string)
		nodes, ok := graph[vertex]
		if ok {
			for _, node := range nodes {
				_, exists := closed[node]
				if !exists {
					stack.Push(node)
					closed[node] = member
				}
			}
			fmt.Println(vertex, " ")
		}
	}
}
