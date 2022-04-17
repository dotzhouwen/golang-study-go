package main

import "fmt"

/**
并查集 684. 冗余连接
*/
func findRedundantConnection(edges [][]int) []int {
	unionFind := NewUnionFind(len(edges) + 1)

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		connected := unionFind.isConnected(u, v)
		if connected {
			return []int{u, v}
		} else {
			unionFind.union(u, v)
		}
	}
	return nil
}

type UnionFind struct {
	parent []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
	}
}

func (u *UnionFind) union(x int, y int) {
	rootX := u.find(x)
	rootY := u.find(y)
	u.parent[rootX] = rootY
}

func (u *UnionFind) find(x int) int {
	for x != u.parent[x] {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *UnionFind) isConnected(x int, y int) bool {
	return u.find(x) == u.find(y)
}

func main() {
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	connection := findRedundantConnection(edges)
	fmt.Println(connection)
}
