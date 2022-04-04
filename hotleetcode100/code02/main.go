package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintNodeList(node *ListNode) {
	cur := node
	builder := strings.Builder{}
	for cur != nil {
		builder.WriteString(fmt.Sprintf("%d", cur.Val))
		if cur.Next != nil {
			builder.WriteString(fmt.Sprintf(" -> "))
		}
		cur = cur.Next
	}
	fmt.Println(builder.String())
}

func createListNode(value int) *ListNode {
	if value == 0 {
		return &ListNode{Val: 0}
	}

	var front *ListNode
	var rear *ListNode

	for value > 0 {
		valNode := &ListNode{Val: value % 10}
		if front == nil {
			front = valNode
			rear = valNode
		} else {
			rear.Next = &ListNode{Val: value % 10}
			rear = rear.Next
		}
		value = value / 10
	}
	return front
}

func reverseListNode(l *ListNode) *ListNode {
	var front *ListNode
	cur := l
	for cur != nil {
		if front == nil {
			front = cur
			cur = cur.Next
			front.Next = nil
		} else {
			curNext := cur.Next
			cur.Next = front
			front = cur
			cur = curNext
		}
	}
	return front
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	rl1 := reverseListNode(l1)
	rl2 := reverseListNode(l2)

	cur1 := rl1
	value1 := 0
	for cur1 != nil {
		value1 = value1*10 + cur1.Val
		cur1 = cur1.Next
	}

	cur2 := rl2
	value2 := 0
	for cur2 != nil {
		value2 = value2*10 + cur2.Val
		cur2 = cur2.Next
	}

	node := createListNode(value1 + value2)

	fmt.Printf("value1:%d, value2:%d\n", value1, value2)

	return node
}

func main() {
	//node11 := ListNode{Val: 2}
	//node12 := ListNode{Val: 4}
	//node13 := ListNode{Val: 3}
	//
	//node21 := ListNode{Val: 5}
	//node22 := ListNode{Val: 6}
	//node23 := ListNode{Val: 4}
	//
	//node11.Next = &node12
	//node12.Next = &node13
	//
	//node21.Next = &node22
	//node22.Next = &node23

	numbers := addTwoNumbers(&ListNode{Val: 0}, &ListNode{Val: 0})

	PrintNodeList(numbers)

}
