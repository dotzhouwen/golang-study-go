package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(data int) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

type BinarySearchTree struct {
	root *Node
}

// MidOrderTree 中序遍历 递归实现
func MidOrderTree(root *Node) {
	if root != nil {
		MidOrderTree(root.Left)
		fmt.Printf("%v ", root.Data)
		MidOrderTree(root.Right)
	}
}

func (tree *BinarySearchTree) Delete(data int) error {
	if tree.root == nil {
		return errors.New("二叉搜索树为空")
	}
	var pp *Node = nil
	p := tree.root

	for p != nil && p.Data != data {
		pp = p
		if data < p.Data {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	if p == nil {
		return errors.New("未查询到结果")
	}

	if p.Left != nil && p.Right != nil {
		minPP := p
		minP := p.Right
		for minP.Left != nil {
			minPP = minP
			minP = minP.Left
		}

		p.Data = minP.Data

		pp = minPP
		p = minP
	}

	var child *Node = nil
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	} else {
		child = nil
	}

	if pp == nil {
		tree.root = nil
	} else {
		if pp.Left == p {
			pp.Left = child
		} else if pp.Right == p {
			pp.Right = child
		}
	}
	return nil
}

func (tree *BinarySearchTree) Find(data int) *Node {
	if tree.root == nil {
		return nil
	}

	p := tree.root
	for p != nil {
		if p.Data == data {
			break
		} else if data < p.Data {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	return p
}

func (tree *BinarySearchTree) Insert(data int) error {
	if tree.root == nil {
		tree.root = NewNode(data)
		return nil
	}

	p := tree.root
	for p != nil {
		if data < p.Data {
			if p.Left == nil {
				p.Left = NewNode(data)
				break
			}
			p = p.Left
		} else if data > p.Data {
			if p.Right == nil {
				p.Right = NewNode(data)
				break
			}
			p = p.Right
		}
	}
	return nil
}

func main() {
	tree := new(BinarySearchTree)
	tree.Insert(9)
	tree.Insert(8)
	tree.Insert(10)
	tree.Insert(7)
	tree.Insert(23)
	tree.Insert(15)

	MidOrderTree(tree.root)
	fmt.Println()

	find := tree.Find(23)
	fmt.Println("find:", find.Data)

	tree.Delete(10)
	MidOrderTree(tree.root)

	fmt.Println()

	tree.Delete(15)
	MidOrderTree(tree.root)
	fmt.Println()
}
