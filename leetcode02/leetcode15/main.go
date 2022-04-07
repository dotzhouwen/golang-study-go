package main

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

/**
N叉树的层序遍历
*/
func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	queue := NewQueue()
	queue.Add(root)

	res := make([][]int, 0)
	for !queue.IsEmpty() {
		size := queue.Len()
		row := make([]int, 0)

		for i := 0; i < size; i++ {
			e := queue.Remove().(*Node)
			if e.Children != nil {
				for _, child := range e.Children {
					queue.Add(child)
				}
			}
			row = append(row, e.Val)
		}
		res = append(res, row)
	}
	return res
}

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	ls := list.New()
	return &Queue{list: ls}
}

func (q *Queue) Add(value interface{}) {
	q.list.PushBack(value)
}

func (q *Queue) Remove() interface{} {
	e := q.list.Front()
	if e != nil {
		q.list.Remove(e)
		return e.Value
	}
	return nil
}

func (q *Queue) IsEmpty() bool {
	return q.list.Len() == 0
}

func (q *Queue) Len() int {
	return q.list.Len()
}

func main() {

}
