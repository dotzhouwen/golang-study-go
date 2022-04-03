package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	if len(preorder) == 1 {
		root.Left = nil
		root.Right = nil
	} else {
		index := 0
		for inorder[index] != rootVal {
			index++
		}
		leftInorder := inorder[0:index]
		rightInorder := inorder[index+1:]

		leftPreorder := preorder[1 : len(leftInorder)+1]
		rightPreorder := preorder[len(leftPreorder)+1:]

		root.Left = buildTree(leftPreorder, leftInorder)
		root.Right = buildTree(rightPreorder, rightInorder)
	}
	return root
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	tree := buildTree(preorder, inorder)
	fmt.Println(tree)
}
