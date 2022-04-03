package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	helper := sumNumbersHelper(root, 0)

	return helper
}

func sumNumbersHelper(root *TreeNode, path int) int {
	if root == nil {
		return 0
	}
	path = path*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return path
	}

	left := sumNumbersHelper(root.Left, path)
	right := sumNumbersHelper(root.Right, path)

	return left + right
}

func createTree() *TreeNode {
	root := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}

	root.Left = &node2
	root.Right = &node3

	return &root
}

func main() {
	tree := createTree()
	numbers := sumNumbers(tree)
	fmt.Println(numbers)
}
