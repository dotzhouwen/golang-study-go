package main

import (
	queue2 "awesomeProject/algorithm/dfs/dfs01/queue"
	"fmt"
)

func main() {

	queue := queue2.NewQueue()
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	for !queue.IsEmpty() {
		fmt.Println(queue.Remove())
	}

}
