package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Key  int
	Next *Node
}

type MyHashSet struct {
	defaultSize int
	nodes       []*Node
}

func PrintNode(node *Node) string {
	builder := strings.Builder{}
	for node != nil {
		builder.WriteString(fmt.Sprintf("{ node: %d }", node.Key))
		if node.Next != nil {
			builder.WriteString(fmt.Sprintf(" -> "))
		}
		node = node.Next
	}
	return builder.String()
}

func (this *MyHashSet) Print() {
	for _, node := range this.nodes {
		if node != nil {
			s := PrintNode(node)
			fmt.Println(s)
		}
	}
}

func Constructor() MyHashSet {
	return MyHashSet{
		defaultSize: 100,
		nodes:       make([]*Node, 100),
	}
}

func (this *MyHashSet) index(key int) int {
	return key % this.defaultSize
}

func (this *MyHashSet) Add(key int) {
	index := this.index(key)
	node := this.nodes[index]

	if node == nil {
		newNode := Node{Key: key, Next: nil}
		this.nodes[index] = &newNode
	} else {
		var preNode *Node
		cur := node
		for cur != nil {
			if cur.Key == key {
				return
			}
			preNode = cur
			cur = cur.Next
		}
		newNode := Node{Key: key, Next: nil}
		preNode.Next = &newNode
	}
}

func (this *MyHashSet) Remove(key int) {
	index := this.index(key)
	node := this.nodes[index]
	if node != nil {
		var preNode *Node
		cur := node
		for cur != nil {
			if cur.Key == key {
				if preNode == nil {
					this.nodes[index] = cur.Next
				} else {
					preNode.Next = cur.Next
				}
			}
			preNode = cur
			cur = cur.Next
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	index := this.index(key)
	node := this.nodes[index]
	if node != nil {
		cur := node
		for cur != nil {
			if key == cur.Key {
				return true
			}
			cur = cur.Next
		}
	}
	return false
}

func main() {
	hashSet := Constructor()
	hashSet.Add(1)
	hashSet.Add(101)
	hashSet.Add(201)
	hashSet.Add(301)

	hashSet.Add(2)
	hashSet.Add(202)
	hashSet.Add(302)

	hashSet.Print()

	fmt.Println(hashSet.Contains(302))
	hashSet.Remove(101)
	hashSet.Print()

	fmt.Println("remove 1")
	hashSet.Remove(1)
	hashSet.Print()

	fmt.Println("remove 201")
	hashSet.Remove(201)
	hashSet.Print()

	fmt.Println("remove 301")
	hashSet.Remove(301)
	hashSet.Print()
}
