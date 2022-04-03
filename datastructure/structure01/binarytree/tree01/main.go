package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := createTree()
	traversal := preorderTraversal(tree)
	fmt.Println(traversal)
}

func createTree() *TreeNode {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}

	node1.Right = &node2
	node2.Left = &node3

	return &node1
}

func preorderTraversal(root *TreeNode) []int {
	collector := make([]int, 0)
	preorderTraversalHelper(root, &collector)
	return collector
}

func preorderTraversalHelper(root *TreeNode, collector *[]int) {
	if root != nil {
		*collector = append(*collector, root.Val)
		preorderTraversalHelper(root.Left, collector)
		preorderTraversalHelper(root.Right, collector)
	}
}
