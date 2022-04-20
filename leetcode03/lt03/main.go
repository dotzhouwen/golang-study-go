package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

/**
克隆图 133. 克隆图 深度优先遍历
*/
func cloneGraph(node *Node) *Node {
	hash := make(map[*Node]*Node, 0)
	return dfs(node, hash)
}

func dfs(node *Node, hash map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if _, exists := hash[node]; exists {
		return hash[node]
	}

	newNode := &Node{Val: node.Val, Neighbors: []*Node{}}
	hash[node] = newNode

	for _, n := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, dfs(n, hash))
	}
	return newNode
}

func createNode() *Node {
	node := &Node{
		Val:       1,
		Neighbors: []*Node{&Node{Val: 2}, &Node{Val: 3}},
	}
	return node
}

func main() {
	node := createNode()
	graph := cloneGraph(node)
	fmt.Println(graph)
}
