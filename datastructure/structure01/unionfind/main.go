package main

import "fmt"

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		count:  n,
		parent: parent,
	}
}

func (union *UnionFind) PrintArr() {
	for _, p := range union.parent {
		fmt.Printf("%d\t", p)
	}
	fmt.Println()
}

func (union *UnionFind) Connected(p int, q int) bool {
	return union.find(p) == union.find(q)
}

func (union *UnionFind) Union(p int, q int) {
	rootP := union.find(p)
	rootQ := union.find(q)
	if rootP == rootQ {
		return
	}
	union.parent[rootP] = rootQ
	union.count--
}

func (union *UnionFind) Count() int {
	return union.count
}

func (union *UnionFind) find(x int) int {
	for union.parent[x] != x {
		x = union.parent[x]
	}
	return x
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
