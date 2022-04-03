package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	rootMount := postorder[len(postorder)-1]
	root := TreeNode{Val: rootMount}

	if len(inorder) == 1 {
		root.Left = nil
		root.Right = nil
	} else {
		index := 0
		for inorder[index] != rootMount {
			index++
		}
		inOrderLeft := inorder[0:index]
		inOrderRight := inorder[index+1 : len(inorder)]

		postOrderLeft := postorder[0:len(inOrderLeft)]
		postOrderRight := postorder[len(inOrderLeft) : len(postorder)-1]

		root.Left = buildTree(inOrderLeft, postOrderLeft)
		root.Right = buildTree(inOrderRight, postOrderRight)
	}
	return &root
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	tree := buildTree(inorder, postorder)
	fmt.Println(tree)
}
