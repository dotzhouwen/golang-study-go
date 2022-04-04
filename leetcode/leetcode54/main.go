package main

import "fmt"

type UnionFind struct {
	parent []int
	size   int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
		size:   n,
		count:  n,
	}
}

func (union *UnionFind) find(x int) int {
	for x != union.parent[x] {
		x = union.parent[x]
	}
	return x
}

func (union *UnionFind) union(p int, q int) {
	pRoot := union.find(p)
	qRoot := union.find(q)
	if pRoot == qRoot {
		return
	}
	union.parent[pRoot] = qRoot
	union.count--
}

func possibleBipartition(n int, dislikes [][]int) bool {

	unionFind := NewUnionFind(n + 1)

	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			canAdd := true
			for _, dislike := range dislikes {
				first, second := dislike[0], dislike[1]
				if first == i && second == j || first == j && second == i {
					canAdd = false
					break
				}
			}
			if canAdd {
				fmt.Printf("i:%d, j:%d\n", i, j)
				unionFind.union(i, j)
			}
		}
	}

	fmt.Println(unionFind.count)

	return false
}

func main() {
	n := 4
	dislikes := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
	}
	possibleBipartition(n, dislikes)
}
