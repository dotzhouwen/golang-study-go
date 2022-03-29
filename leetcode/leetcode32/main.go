package main

import (
	queue2 "awesomeProject/algorithm/dfs/dfs01/queue"
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

func CreateSymmetricTree() *TreeNode {
	treeNode1 := TreeNode{Val: 1}
	treeNode2_1 := TreeNode{Val: 2}
	treeNode2_2 := TreeNode{Val: 2}
	treeNode3_1 := TreeNode{Val: 3}
	treeNode3_2 := TreeNode{Val: 3}
	treeNode4_1 := TreeNode{Val: 4}
	treeNode4_2 := TreeNode{Val: 4}

	treeNode1.Left = &treeNode2_1
	treeNode1.Right = &treeNode2_2

	treeNode2_1.Left = &treeNode4_1
	treeNode2_1.Right = &treeNode3_1

	treeNode2_2.Left = &treeNode3_2
	treeNode2_2.Right = &treeNode4_2

	return &treeNode1
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

/**
对称二叉树
							1
					2					2
			  	4		  3		   3		 4

[4 2 3 1 3 2 4]

*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}

func TraverseTree3(root *TreeNode) [][]int {
	queue := queue2.NewQueue()
	queue.Add(root)
	result := make([][]int, 0)

	for !queue.IsEmpty() {
		levelNum := queue.Len()
		levelSlice := make([]int, levelNum, levelNum)

		for i := 0; i < levelNum; i++ {
			e := queue.Remove().(*TreeNode)
			if e.Left != nil {
				queue.Add(e.Left)
			}
			if e.Right != nil {
				queue.Add(e.Right)
			}
			levelSlice[i] = e.Val
		}

		result = append(result, levelSlice)
	}
	return result
}

func TraverseTree2(root *TreeNode) {
	queue := queue2.NewQueue()
	queue.Add(root)
	for !queue.IsEmpty() {
		e := queue.Remove().(*TreeNode)
		if e.Left != nil {
			queue.Add(e.Left)
		}
		if e.Right != nil {
			queue.Add(e.Right)
		}

		fmt.Println(e.Val)
	}
}

func levelOrder(root *TreeNode) [][]int {

	return nil
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
	tree := CreateSymmetricTree()
	tree3 := TraverseTree3(tree)

	for _, slice := range tree3 {
		fmt.Println(slice)
	}
}
