package main

import (
	stack2 "awesomeProject/algorithm/dfs/dfs01/stack"
	"awesomeProject/leetcode/util"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
二叉树的前序遍历 非递归写法
*/
func preorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := stack2.NewStack()
	node := root

	for !stack.Empty() || node != nil {
		for node != nil {
			res = append(res, node.Val)
			stack.Push(node)
			node = node.Left
		}
		node = stack.Pop().(*TreeNode)
		node = node.Right
	}

	return res
}

/**
中序遍历，非递归
*/
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := util.NewStack()
	p := root
	for p != nil {
		stack.Push(p)
		p = p.Left
	}

	for !stack.Empty() {
		node := stack.Pop().(*TreeNode)
		res = append(res, node.Val)
		p = node.Right

		for p != nil {
			stack.Push(p)
			p = p.Left
		}
	}

	return res
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	stack := stack2.NewStack()
	stack.Push(root)

	for !stack.Empty() {
		e := stack.Pop().(*TreeNode)
		ans = append(ans, e.Val)

		if e.Right != nil {
			stack.Push(e.Right)
		}
		if e.Left != nil {
			stack.Push(e.Left)
		}
	}

	return ans
}

func CreateTree() *TreeNode {
	root := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}

	node4 := TreeNode{Val: 4}
	node5 := TreeNode{Val: 5}
	node6 := TreeNode{Val: 6}

	root.Left = &node2
	root.Right = &node3

	node2.Right = &node4
	node3.Left = &node5
	node3.Right = &node6

	return &root
}

func recoverFromPreorder(traversal string) *TreeNode {

	return nil
}

func main() {
	tree := CreateTree()
	traversal := inorderTraversal(tree)
	fmt.Println(traversal)
}
