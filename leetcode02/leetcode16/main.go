package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
993. 二叉树的堂兄弟节点
*/
func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)

		var xFindNodeParent *TreeNode
		var yFindNodeParent *TreeNode

		for i := 0; i < size; i++ {
			e := queue[0]
			queue = queue[1:]

			if e.Left != nil {
				queue = append(queue, e.Left)
				if e.Left.Val == x {
					xFindNodeParent = e
				}
				if e.Left.Val == y {
					yFindNodeParent = e
				}
			}
			if e.Right != nil {
				queue = append(queue, e.Right)
				if e.Right.Val == x {
					xFindNodeParent = e
				}
				if e.Right.Val == y {
					yFindNodeParent = e
				}
			}
		}
		if xFindNodeParent != nil && yFindNodeParent != nil {
			if xFindNodeParent != yFindNodeParent {
				return true
			} else {
				return false
			}
		}
	}

	return false
}

func creatTree() *TreeNode {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	node5 := TreeNode{Val: 5}

	node1.Left = &node2
	node1.Right = &node3

	node2.Left = &node4
	node3.Right = &node5

	return &node1
}

func main() {
	tree := creatTree()
	cousins := isCousins(tree, 4, 5)
	fmt.Println(cousins)
}
