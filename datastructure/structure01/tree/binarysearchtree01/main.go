package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (node *Node) GetData() string {
	return fmt.Sprintf("%v", node.Data)
}

type BinarySearchTree struct {
	root *Node
}

func NewBinarySearchTree(node *Node) *BinarySearchTree {
	return &BinarySearchTree{
		root: node,
	}
}

func midOrderTraverse(root *Node) {
	if root == nil {
		return
	}
	midOrderTraverse(root.Left)
	fmt.Printf("%v ", root.GetData())
	midOrderTraverse(root.Right)
}

func (tree *BinarySearchTree) Delete(data int) error {
	if tree.root == nil {
		return errors.New("二叉树为空")
	}
	p := tree.root
	var pp *Node = nil

	for p != nil && p.Data.(int) != data {
		pp = p
		if data > p.Data.(int) {
			p = p.Right
		} else {
			p = p.Left
		}
	}

	if p == nil {
		return errors.New("待删除结点不存在")
	}

	if p.Left != nil && p.Right != nil {
		minP := p.Right
		minPP := p
		for minP.Left != nil {
			minPP = minP
			minP = minP.Left
		}
		p.Data = minP.Data
		p = minP
		pp = minPP
	}

	var child *Node
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	} else {
		child = nil
	}

	if pp == nil {
		tree.root = nil
	} else if pp.Left == p {
		pp.Left = child
	} else if pp.Right == p {
		pp.Right = child
	}
	return nil
}

func (tree *BinarySearchTree) Insert(data interface{}) error {
	var value int
	var ok bool
	if value, ok = data.(int); !ok {
		return errors.New("暂时只支持int类型数据")
	}

	if node := tree.root; node == nil {
		tree.root = NewNode(value)
		return nil
	} else {
		for node != nil {
			if value < node.Data.(int) {
				if node.Left == nil {
					node.Left = NewNode(value)
					break
				}
				node = node.Left
			} else if value > node.Data.(int) {
				if node.Right == nil {
					node.Right = NewNode(value)
					break
				}
				node = node.Right
			} else {
				return errors.New("对应数据已经存在")
			}
		}
		return nil
	}
}

func (tree *BinarySearchTree) Find(data int) *Node {
	node := tree.root

	if node == nil {
		return nil
	}
	for node != nil {
		if data < node.Data.(int) {
			node = node.Left
		} else if data > node.Data.(int) {
			node = node.Right
		} else {
			return node
		}
	}
	return nil
}

func main() {
	tree := NewBinarySearchTree(nil)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(4)

	fmt.Println("中序遍历二叉树")
	midOrderTraverse(tree.root)
	fmt.Println()

	fmt.Println("查找值为3的节点")
	node := tree.Find(3)
	fmt.Printf("%v\n", node)

	tree.Delete(4)
	midOrderTraverse(tree.root)
	fmt.Println()

	tree.Delete(3)
	midOrderTraverse(tree.root)
	fmt.Println()

	tree.Delete(1)
	midOrderTraverse(tree.root)
	fmt.Println()
}
