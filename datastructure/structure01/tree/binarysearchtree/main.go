package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

type BinarySearchTree struct {
	tree *Node
}

func (b *BinarySearchTree) InOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	b.InOrderTraversal(root.left)
	fmt.Printf("%v ", root.data)
	b.InOrderTraversal(root.right)
}

func (b *BinarySearchTree) Find(value int) *Node {
	node := b.tree
	for node != nil {
		if value < node.data {
			node = node.left
		} else if value > node.data {
			node = node.right
		} else {
			return node
		}
	}
	return nil
}

func (b *BinarySearchTree) Insert(value int) bool {
	if b.tree == nil {
		b.tree = NewNode(value)
		return true
	}
	p := b.tree
	for p != nil {
		if value > p.data {
			if p.right == nil {
				p.right = NewNode(value)
				return true
			}
			p = p.right
		} else {
			if p.left == nil {
				p.left = NewNode(value)
				return true
			}
			p = p.left
		}
	}
	return false
}

func main() {
	b := new(BinarySearchTree)
	b.Insert(13)
	b.Insert(8)
	b.Insert(18)
	b.Insert(6)
	b.Insert(10)
	b.Insert(16)
	b.Insert(20)

	b.InOrderTraversal(b.tree)

	fmt.Println()
	node := b.Find(16)
	fmt.Println(node.data)
}
