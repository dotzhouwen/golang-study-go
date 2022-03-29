package queue

import (
	"container/list"
)

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
