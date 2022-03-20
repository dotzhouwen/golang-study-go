package main

import "fmt"

type StackNode struct {
	data int
	next *StackNode
}

type LinkStack struct {
	top   *StackNode
	count int
}

func NewStackNode(data int) *StackNode {
	return &StackNode{
		data: data,
		next: nil,
	}
}

func (s *LinkStack) Push(data int) bool {
	node := NewStackNode(data)
	node.next = s.top
	s.top = node
	s.count++
	return true
}

func (s *LinkStack) Pop() *StackNode {
	if s.top == nil {
		return nil
	}
	node := s.top
	s.top = s.top.next
	s.count--
	return node
}

func main() {
	stack := new(LinkStack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
