package main

import (
	queue2 "awesomeProject/algorithm/dfs/dfs01/queue"
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func levelTraversal(root *Node) {
	queue := queue2.NewQueue()
	queue.Add(root)

	for !queue.IsEmpty() {
		levelNum := queue.Len()

		var preNode *Node
		for i := 0; i < levelNum; i++ {
			e := queue.Remove().(*Node)

			if i == levelNum-1 {
				e.Next = nil
			}

			if preNode != nil {
				preNode.Next = e
			}

			preNode = e

			if e.Left != nil {
				queue.Add(e.Left)
			}
			if e.Right != nil {
				queue.Add(e.Right)
			}
			fmt.Printf("%d ", e.Val)
		}
		fmt.Println()
	}
}

/**
填充每个节点的下一个右侧节点指针
*/
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := queue2.NewQueue()

	root.Next = nil
	queue.Add(root)

	for !queue.IsEmpty() {
		levelNum := queue.Len()
		var preNode *Node
		for i := 0; i < levelNum; i++ {
			e := queue.Remove().(*Node)

			if i == levelNum-1 {
				e.Next = nil
			}

			if preNode != nil {
				preNode.Next = e
			}

			preNode = e

			if e.Left != nil {
				queue.Add(e.Left)
			}
			if e.Right != nil {
				queue.Add(e.Right)
			}
		}
	}
	return root
}

func createTree() *Node {
	root := Node{Val: 1}
	node2 := Node{Val: 2}
	node3 := Node{Val: 3}
	node4 := Node{Val: 4}
	node5 := Node{Val: 5}
	node6 := Node{Val: 6}
	node7 := Node{Val: 7}

	root.Left = &node2
	root.Right = &node3

	node2.Left = &node4
	node2.Right = &node5

	node3.Left = &node6
	node3.Right = &node7

	return &root
}

func main() {
	tree := createTree()
	levelTraversal(tree)

	fmt.Println(tree)
}
