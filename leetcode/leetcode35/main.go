package main

import (
	"awesomeProject/leetcode/util"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := util.NewQueue()
	queue.Add(root)

	res := make([][]int, 0)
	for !queue.IsEmpty() {
		levelNum := queue.Len()
		levelSlice := make([]int, levelNum, levelNum)

		for i := 0; i < levelNum; i++ {
			e := queue.Remove().(*TreeNode)
			levelSlice[i] = e.Val

			if e.Left != nil {
				queue.Add(e.Left)
			}
			if e.Right != nil {
				queue.Add(e.Right)
			}
		}
		res = append(res, levelSlice)
	}
	return res
}

func preOrderTraversalRecursive(root *TreeNode) []int {
	collector := make([]int, 0)
	preOrderTraversalRecursiveHelper(root, &collector)
	return collector
}

func preOrderTraversalRecursiveHelper(root *TreeNode, coll *[]int) {
	if root != nil {
		*coll = append(*coll, root.Val)
		preOrderTraversalRecursiveHelper(root.Left, coll)
		preOrderTraversalRecursiveHelper(root.Right, coll)
	}
}

func preOrderTraversal(root *TreeNode) []int {
	stack := util.NewStack()
	result := make([]int, 0)

	stack.Push(root)
	for !stack.Empty() {
		e := stack.Pop().(*TreeNode)
		if e.Right != nil {
			stack.Push(e.Right)
		}
		if e.Left != nil {
			stack.Push(e.Left)
		}
		result = append(result, e.Val)
	}
	return result
}

func inorderTraversal2(root *TreeNode) []int {

	return nil
}

func postorderTraversal(root *TreeNode) []int {
	collector := make([]int, 0)
	postorderTraversalHelper(root, &collector)
	return collector
}

func postorderTraversalHelper(root *TreeNode, collector *[]int) {
	if root != nil {
		postorderTraversalHelper(root.Left, collector)
		postorderTraversalHelper(root.Right, collector)
		*collector = append(*collector, root.Val)
	}
}

func inorderTraversal(root *TreeNode) []int {
	collector := make([]int, 0)
	inorderTraversalHelper(root, &collector)
	return collector
}

func inorderTraversalHelper(root *TreeNode, collector *[]int) {
	if root != nil {
		inorderTraversalHelper(root.Left, collector)
		*collector = append(*collector, root.Val)
		inorderTraversalHelper(root.Right, collector)
	}
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

func main() {
	tree := CreateSymmetricTree()
	traversal := levelOrder(tree)
	fmt.Println(traversal)
}
