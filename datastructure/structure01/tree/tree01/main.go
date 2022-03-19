package main

import "fmt"

type TreeNode struct {
	data  string
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(data string) *TreeNode {
	return &TreeNode{
		data: data,
	}
}

var i = -1

func CreateBiTree(arr []string) *TreeNode {
	i = i + 1
	if i >= len(arr) {
		return nil
	}
	var node *TreeNode

	if arr[i] == "#" {
		return nil
	} else {
		node = NewTreeNode(arr[i])
		node.left = CreateBiTree(arr)
		node.right = CreateBiTree(arr)
	}
	return node
}

func main() {
	var treeIds = []string{"A", "B", "#", "#", "C", "D", "#", "#", "#"}
	//t := CreateBiTree(treeIds)
	fmt.Println(treeIds)
}
