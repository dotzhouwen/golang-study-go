package main

import (
	stack2 "awesomeProject/algorithm/dfs/dfs01/stack"
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewBSTTree() *TreeNode {
	node2 := TreeNode{Val: 2}
	node1 := TreeNode{Val: 1}
	node3 := TreeNode{Val: 3}

	node2.Left = &node1
	node2.Right = &node3

	return &node2
}

func NewTree() *TreeNode {
	node3 := TreeNode{Val: 3}
	node9 := TreeNode{Val: 9}
	node20 := TreeNode{Val: 20}
	node15 := TreeNode{Val: 15}
	node7 := TreeNode{Val: 7}
	node31 := TreeNode{Val: 31}
	node33 := TreeNode{Val: 33}

	node7.Left = &node31
	node7.Right = &node33

	node3.Left = &node9
	node3.Right = &node20

	node20.Left = &node15
	node20.Right = &node7

	return &node3
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Val >= root.Val || root.Right != nil && root.Right.Val <= root.Val {
		return false
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}

func TraverseTree(root *TreeNode) {
	stack := stack2.NewStack()
	stack.Push(root)
	for !stack.Empty() {
		pop := stack.Pop().(*TreeNode)

		if pop.Right != nil {
			stack.Push(pop.Right)
		}
		if pop.Left != nil {
			stack.Push(pop.Left)
		}

		fmt.Println("pop:", pop.Val)
	}
}

func main() {
	tree := NewBSTTree()
	bst := isValidBST(tree)
	fmt.Println(bst)
}
