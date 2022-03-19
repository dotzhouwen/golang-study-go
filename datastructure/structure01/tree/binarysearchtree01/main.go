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

	fmt.Println("查找值为3的节点")
	node := tree.Find(3)
	fmt.Printf("%v\n", node)
}
