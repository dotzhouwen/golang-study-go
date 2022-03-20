package main

import "fmt"

// BinarySearchTreeNode 二叉查找树节点
type BinarySearchTreeNode struct {
	Value int64
	Times int64
	Left  *BinarySearchTreeNode
	Right *BinarySearchTreeNode
}

type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		Root: nil,
	}
}

func (tree *BinarySearchTree) MidOrder() {
	tree.Root.MidOrder()
}

func (node *BinarySearchTreeNode) MidOrder() {
	if node == nil {
		return
	}
	node.Left.MidOrder()
	for i := 0; i <= int(node.Times); i++ {
		fmt.Printf("%v ", node.Value)
	}
	node.Right.MidOrder()
}

func (tree *BinarySearchTree) Add(value int64) {
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{Value: value}
		return
	}
	tree.Root.Add(value)
}

func (node *BinarySearchTreeNode) Add(value int64) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &BinarySearchTreeNode{Value: value}
		} else {
			node.Left.Add(value)
		}
	} else if value > node.Value {
		if node.Right == nil {
			node.Right = &BinarySearchTreeNode{Value: value}
		} else {
			node.Right.Add(value)
		}
	} else {
		node.Times = node.Times + 1
	}
}

func (tree *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMinValue()
}

func (node *BinarySearchTreeNode) FindMinValue() *BinarySearchTreeNode {
	if node.Left == nil {
		return node
	}
	return node.Left.FindMinValue()
}

func (tree *BinarySearchTree) FindMaxValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMaxValue()
}

func (node *BinarySearchTreeNode) FindMaxValue() *BinarySearchTreeNode {
	if node.Right == nil {
		return node
	}
	return node.Right.FindMaxValue()
}

func (tree *BinarySearchTree) Find(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.Find(value)
}

func (node *BinarySearchTreeNode) Find(value int64) *BinarySearchTreeNode {
	if node == nil {
		return nil
	}
	if node.Value == value {
		return node
	} else if value < node.Value {
		return node.Left.Find(value)
	} else {
		return node.Right.Find(value)
	}
}

func (node *BinarySearchTreeNode) FindParent(value int64) *BinarySearchTreeNode {
	if value < node.Value {
		leftTree := node.Left
		if leftTree == nil {
			return nil
		}
		if leftTree.Value == value {
			return node
		} else {
			return leftTree.FindParent(value)
		}
	} else {
		rightTree := node.Right
		if rightTree == nil {
			return nil
		}
		if rightTree.Value == value {
			return node
		} else {
			return rightTree.FindParent(value)
		}
	}
}

func (tree *BinarySearchTree) Delete(value int64) {
	if tree.Root == nil {
		return
	}
	node := tree.Find(value)
	if node == nil {
		return
	}

	parent := tree.Root.FindParent(value)
	if parent == nil && node.Left == nil && node.Right == nil {
		tree.Root = nil
		return
	} else if parent != nil && node.Left == nil && node.Right == nil {
		if parent.Left != nil && value == parent.Left.Value {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	} else if node.Left != nil && node.Right != nil {
		minNode := node.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		tree.Delete(minNode.Value)
		node.Value = minNode.Value
		node.Times = minNode.Times
	} else {
		if parent == nil {
			if node.Left != nil {
				tree.Root = node.Left
			} else {
				tree.Root = node.Right
			}
			return
		}
		if node.Left != nil {
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Left
			} else {
				parent.Right = node.Right
			}
		} else {
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		}
	}
}

func main() {
	values := []int64{3, 6, 8, 20, 9, 2, 6, 8, 9, 3, 5, 40, 7, 9, 13, 6, 8}
	tree := NewBinarySearchTree()
	for _, v := range values {
		tree.Add(v)
	}

	tree.MidOrder()

	minValue := tree.FindMinValue()
	maxValue := tree.FindMaxValue()

	fmt.Println()
	fmt.Printf("minValue: %d, maxValue:%d\n", minValue.Value, maxValue.Value)
}
