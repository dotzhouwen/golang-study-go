package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateNewTree() *TreeNode {
	node5 := TreeNode{Val: 5}
	node4 := TreeNode{Val: 4}
	node8 := TreeNode{Val: 8}
	node11 := TreeNode{Val: 11}
	node13 := TreeNode{Val: 13}
	node42 := TreeNode{Val: 4}

	node5.Left = &node4
	node5.Right = &node8

	node4.Left = &node11
	node8.Left = &node13
	node8.Right = &node42

	return &node5
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root != nil {
		if root.Left == nil && root.Right == nil && root.Val == targetSum {
			return true
		}
		return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
	}
	return false
}

func main() {
	tree := CreateNewTree()
	sum := hasPathSum(tree, 13)
	fmt.Println(sum)
}
