package main

import "fmt"

type SqStack struct {
	SElemType []interface{}
	top       int
	cap       int
}

func NewSqStack(cap int) *SqStack {
	return &SqStack{
		SElemType: make([]interface{}, cap, cap),
		top:       -1,
		cap:       cap,
	}
}

func (s *SqStack) Push(e interface{}) bool {
	if s.top == s.cap-1 {
		return false
	}
	s.top++
	s.SElemType[s.top] = e

	return true
}

func (s *SqStack) Pop() interface{} {
	if s.top == -1 {
		return nil
	}
	e := s.SElemType[s.top]
	s.top--
	return e
}

func main() {
	s := NewSqStack(3)
	fmt.Println(s.Push(1))
	fmt.Println(s.Push(2))
	fmt.Println(s.Push(3))
	fmt.Println(s.Push(4))

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
