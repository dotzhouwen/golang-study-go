package main

import "fmt"

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

/**
通过并查集解决 547. 省份数量
*/
func findCircleNum(isConnected [][]int) int {
	length := len(isConnected)
	unionFind := NewUnionFind(length)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if isConnected[i][j] == 1 {
				unionFind.union(i, j)
			}
		}
	}

	hash := make(map[int]int, 0)
	for i := 0; i < length; i++ {
		root := unionFind.find(i)
		hash[root]++
	}
	return len(hash)
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
