package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
二叉树遍历
*/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	reverse := false
	res := make([][]int, 0)

	for len(queue) > 0 {
		size := len(queue)
		levelSlice := make([]int, 0)
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
		if reverse {
			reverseSlice(levelSlice)
		}
		res = append(res, levelSlice)
		reverse = !reverse
	}
	return res
}

func reverseSlice(slice []int) {
	left := 0
	right := len(slice) - 1
	for left < right {
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}
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
