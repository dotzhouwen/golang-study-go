package main

import "fmt"

type UnionFind struct {
	root []int
	rank []int
}

func NewUnionFind(n int) *UnionFind {
	root := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
		rank[i] = 1
	}
	return &UnionFind{
		root: root,
		rank: rank,
	}
}

func (union *UnionFind) find(x int) int {
	if x == union.root[x] {
		return x
	} else {
		union.root[x] = union.find(union.root[x])
		return union.root[x]
	}
}

func (union *UnionFind) union(p int, q int) {
	rootP := union.find(p)
	rootQ := union.find(q)

	if rootP != rootQ {
		if union.rank[rootP] > union.rank[rootQ] {
			union.root[rootQ] = rootP
		} else if union.rank[rootP] < union.rank[rootQ] {
			union.root[rootP] = rootQ
		} else {
			union.root[rootP] = rootQ
			union.rank[rootQ] = union.rank[rootQ] + 1
		}
	}
}

func (union *UnionFind) isConnected(p int, q int) bool {
	return union.find(p) == union.find(q)
}

func findRedundantConnection(edges [][]int) []int {
	unionFind := NewUnionFind(len(edges) * len(edges[0]))
	var circleEdge []int
	for _, edge := range edges {
		source := edge[0]
		target := edge[1]
		if unionFind.isConnected(source, target) {
			circleEdge = edge
		} else {
			unionFind.union(source, target)
		}
	}

	return circleEdge
}

func main() {
	edges := [][]int{
		{1, 2},
		{1, 3},
		{3, 3},
	}
	connection := findRedundantConnection(edges)
	fmt.Println(connection)
}
