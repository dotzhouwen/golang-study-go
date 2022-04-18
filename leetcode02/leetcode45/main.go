package main

import "fmt"

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

/**
765. 情侣牵手
*/
func minSwapsCouples(row []int) int {
	m := len(row)

	unionFind := NewUnionFind(m / 2)
	for i := 0; i < m; i += 2 {
		unionFind.union(row[i]/2, row[i+1]/2)
	}

	return m/2 - unionFind.getSize()
}

func main() {
	row := []int{5, 4, 2, 6, 3, 1, 0, 7}
	couples := minSwapsCouples(row)
	fmt.Println(couples)
}
