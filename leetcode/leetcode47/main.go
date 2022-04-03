package main

import "fmt"

type UnionFind struct {
	parent []int
	size   int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
		size:   n,
	}
}

func (union *UnionFind) find(u int) int {
	if u == union.parent[u] {
		return u
	} else {
		p := union.find(union.parent[u])
		union.parent[u] = p
		return p
	}
}

func (union *UnionFind) Union(p int, q int) {
	rootP := union.find(p)
	rootQ := union.find(q)
	if rootP == rootQ {
		return
	}
	union.parent[rootP] = rootQ
}

/**
547. 省份数量
*/
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	unionFind := NewUnionFind(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				unionFind.Union(i, j)
			}
		}
	}

	rootCount := make(map[int]struct{})
	for i := 0; i < n; i++ {
		rootCount[unionFind.find(i)] = struct{}{}
	}

	return len(rootCount)
}

func main() {
	isConnected := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	num := findCircleNum(isConnected)
	fmt.Println(num)
}
