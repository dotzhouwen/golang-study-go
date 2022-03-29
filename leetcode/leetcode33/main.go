package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

// TraverseTree3 /**
func TraverseTree3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := NewQueue()
	queue.Add(root)
	result := make([][]int, 0)

	for !queue.IsEmpty() {
		levelNum := queue.Len()
		levelSlice := make([]int, levelNum, levelNum)

		for i := 0; i < levelNum; i++ {
			e := queue.Remove().(*TreeNode)
			if e.Left != nil {
				queue.Add(e.Left)
			}
			if e.Right != nil {
				queue.Add(e.Right)
			}
			levelSlice[i] = e.Val
		}

		result = append(result, levelSlice)
	}
	return result
}

func main() {

}
