package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) String() string {
	return fmt.Sprintf("{ %d left: %v, right:%v }", node.Val, node.Left, node.Right)
}

/**
寻找重复的子树
*/
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	ans := make([]*TreeNode, 0)
	hash := make(map[string]int, 0)
	serializeTree(root, &ans, &hash)
	return ans
}

func serializeTree(node *TreeNode, ans *[]*TreeNode, hash *map[string]int) string {
	if node == nil {
		return "#"
	}
	serial := strconv.Itoa(node.Val) + "," + serializeTree(node.Left, ans, hash) + "," + serializeTree(node.Right, ans, hash)
	(*hash)[serial]++
	if (*hash)[serial] == 2 {
		*ans = append(*ans, node)
	}
	return serial
}

func main() {
	tree := createTree()
	subtrees := findDuplicateSubtrees(tree)
	fmt.Println(subtrees)
}

func createTree() *TreeNode {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 2}
	node4 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 3}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node3.Left = node5

	return node1
}
