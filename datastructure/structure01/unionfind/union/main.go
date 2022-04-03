package main

import "fmt"

type UnionFind struct {
	count  int
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &UnionFind{
		count:  n,
		parent: parent,
		size:   size,
	}
}

func (union *UnionFind) PrintArr() {
	for _, p := range union.parent {
		fmt.Printf("%d\t", p)
	}
	fmt.Println()
}

func (union *UnionFind) find(x int) int {
	for x != union.parent[x] {
		union.parent[x] = union.parent[union.parent[x]]
		x = union.parent[x]
	}
	return x
}

func (union *UnionFind) Union(p int, q int) {
	rootP := union.find(p)
	rootQ := union.find(q)
	if rootP == rootQ {
		return
	}
	if union.size[rootP] > union.size[rootQ] {
		union.parent[rootQ] = rootP
		union.size[rootP] += union.size[rootQ]
	} else {
		union.parent[rootP] = rootQ
		union.size[rootQ] += union.size[rootP]
	}
	union.count--
}

func (union *UnionFind) Connected(p int, q int) bool {
	return union.find(p) == union.find(q)
}

func (union *UnionFind) Count() int {
	return union.count
}

func main() {
	fmt.Println("初始化：")
	unionFind := NewUnionFind(10)
	unionFind.PrintArr()

	fmt.Println()
	fmt.Println("连接5， 6")
	unionFind.Union(5, 6)
	unionFind.PrintArr()
	fmt.Println("测试5 6 是否连接:", unionFind.Connected(5, 6))
	fmt.Println("连通分量:", unionFind.Count())

	fmt.Println()
	fmt.Println("连接2， 3")
	unionFind.Union(2, 3)
	unionFind.PrintArr()
	fmt.Println("测试2 3 是否连接:", unionFind.Connected(2, 3))
	fmt.Println("连通分量:", unionFind.Count())

	fmt.Println()
	fmt.Println("连接3， 5")
	unionFind.Union(3, 5)
	unionFind.PrintArr()
	fmt.Println("测试3 5 是否连接:", unionFind.Connected(3, 5))
	fmt.Println("测试2 6 是否连接:", unionFind.Connected(2, 6))
	fmt.Println("连通分量:", unionFind.Count())
}
