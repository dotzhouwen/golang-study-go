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
等式方程的可满足性
*/
func equationsPossible(equations []string) bool {
	unionFind := NewUnionFind(26)
	for _, equation := range equations {
		chars := []rune(equation)
		if chars[1] == '=' {
			i := int(chars[0] - 'a')
			j := int(chars[3] - 'a')
			unionFind.union(i, j)
		}
	}

	for _, equation := range equations {
		chars := []rune(equation)
		if chars[1] == '!' {
			i := int(chars[0] - 'a')
			j := int(chars[3] - 'a')
			if unionFind.isConnected(i, j) {
				return false
			}
		}
	}
	return true
}

func main() {
	equations := []string{"a==b", "b==c", "a!=c"}
	possible := equationsPossible(equations)
	fmt.Println(possible)
}
