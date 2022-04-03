package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
中序遍历和后序遍历构造二叉树
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}

	if len(postorder) == 1 {
		root.Left = nil
		root.Right = nil
	} else {
		index := 0
		for inorder[index] != postorder[len(postorder)-1] {
			index++
		}

		leftInorder := inorder[0:index]
		rightInorder := inorder[index+1 : len(inorder)]

		leftPostOrder := postorder[0:len(leftInorder)]
		rightPostOrder := postorder[len(leftPostOrder) : len(postorder)-1]

		root.Left = buildTree(leftInorder, leftPostOrder)
		root.Right = buildTree(rightInorder, rightPostOrder)
	}
	return root
}

func main() {
	inorder := []int{2, 1}
	postorder := []int{2, 1}

	tree := buildTree(inorder, postorder)

	fmt.Println(tree)
}
