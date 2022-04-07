package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
二叉树的层序遍历
*/
func levelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)

	if root == nil {
		return nil
	}
	queue = append(queue, root)
	res := make([][]int, 0)
	for len(queue) > 0 {
		levelSlice := make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			e := queue[0]
			queue = queue[1:]
			if e.Left != nil {
				queue = append(queue, e.Left)
			}
			if e.Right != nil {
				queue = append(queue, e.Right)
			}
			levelSlice = append(levelSlice, e.Val)
		}
		res = append(res, levelSlice)
	}
	return res
}

func createTree() *TreeNode {
	node3 := TreeNode{Val: 3}
	node9 := TreeNode{Val: 9}
	node20 := TreeNode{Val: 20}
	node15 := TreeNode{Val: 15}
	node7 := TreeNode{Val: 7}

	node3.Left = &node9
	node3.Right = &node20
	node20.Left = &node15
	node20.Right = &node7

	return &node3
}

func main() {
	tree := createTree()
	order := levelOrder(tree)
	fmt.Println(order)
}
