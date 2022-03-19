package main

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	MaxSize int
	Array   [4]int
	Front   int
	Rear    int
}

func NewCircleQueue() *CircleQueue {
	return &CircleQueue{
		MaxSize: 4,
		Front:   0,
		Rear:    0,
	}
}

func (q *CircleQueue) Push(val int) (err error) {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	q.Array[q.Rear] = val
	q.Rear = (q.Rear + 1) % q.MaxSize
	return nil
}

func (q *CircleQueue) Pop() (val int, err error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	val = q.Array[q.Front]
	q.Front = (q.Front + 1) % q.MaxSize
	return val, nil
}

func (q *CircleQueue) IsFull() bool {
	return (q.Rear+1)%q.MaxSize == q.Front
}

func (q *CircleQueue) IsEmpty() bool {
	return q.Front == q.Rear
}

func (q *CircleQueue) Size() int {
	return (q.Rear - q.Front + q.MaxSize) % q.MaxSize
}

func (q *CircleQueue) Show() {
	size := q.Size()
	tmpFront := q.Front
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%v ", tmpFront, q.Array[tmpFront])
		tmpFront = (tmpFront + 1) % q.MaxSize
	}
	fmt.Println()
}

func main() {
	q := NewCircleQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	err := q.Push(4)
	if err != nil {
		fmt.Println(err.Error())
	}
	q.Show()
	q.Pop()

	q.Show()
}
