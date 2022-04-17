package main

import "fmt"

/**
1319. 连通网络的操作次数
*/
func makeConnected(n int, connections [][]int) int {
	unionFind := NewUnionFind(n)

	count := 0
	for _, connection := range connections {
		u := connection[0]
		v := connection[1]
		if unionFind.isConnected(u, v) {
			count++
		} else {
			unionFind.union(u, v)
		}
	}

	if unionFind.getSize() == 0 {
		return 0
	}

	if count >= unionFind.getSize()-1 {
		return unionFind.getSize() - 1
	} else {
		return -1
	}
}

type UnionFind struct {
	parent []int
	size   int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
		size:   size,
	}
}

func (u *UnionFind) find(x int) int {
	for x != u.parent[x] {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *UnionFind) union(x int, y int) {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX != rootY {
		u.parent[rootX] = rootY
		u.size--
	}
}

func (u *UnionFind) isConnected(x int, y int) bool {
	return u.find(x) == u.find(y)
}

func (u *UnionFind) getSize() int {
	return u.size
}

func main() {
	n := 4
	connections := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
	}
	connected := makeConnected(n, connections)
	fmt.Println(connected)
}
