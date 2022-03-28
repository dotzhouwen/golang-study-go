package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkList struct {
	Front *ListNode
	Rear  *ListNode
}

func NewLinkList() *LinkList {
	return new(LinkList)
}

func (l *LinkList) Print() {
	builder := strings.Builder{}
	cur := l.Front
	for cur != nil {
		builder.WriteString(fmt.Sprintf("%d", cur.Val))
		cur = cur.Next
		if cur != nil {
			builder.WriteString(" -> ")
		}
	}
	fmt.Println(builder.String())
}

func PrintNodeList(node *ListNode) {
	builder := strings.Builder{}
	cur := node
	for cur != nil {
		builder.WriteString(fmt.Sprintf("%d", cur.Val))
		cur = cur.Next
		if cur != nil {
			builder.WriteString(" -> ")
		}
	}
	fmt.Println(builder.String())
}

func CreateLinkList(arr []int) *ListNode {
	linkList := NewLinkList()
	for i := 0; i < len(arr); i++ {
		linkList.PushBack(arr[i])
	}
	return linkList.Front
}

func (l *LinkList) PushBack(val int) {
	if l.Front == nil {
		node := &ListNode{
			Val:  val,
			Next: nil,
		}
		l.Front = node
		l.Rear = node
	} else {
		node := &ListNode{
			Val:  val,
			Next: nil,
		}
		l.Rear.Next = node
		l.Rear = node
	}
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head *ListNode
	if list1.Val < list2.Val {
		head = list1
		head.Next = mergeTwoLists2(list1.Next, list2)
	} else {
		head = list2
		head.Next = mergeTwoLists2(list1, list2.Next)
	}
	return head
}

func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head

	for slow != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

func hasCycle(head *ListNode) bool {
	pointerMap := make(map[*ListNode]int)

	cur := head
	for cur != nil {
		if cur.Next != nil {
			pointerMap[cur.Next]++
			if pointerMap[cur.Next] > 1 {
				return true
			}
		}
		cur = cur.Next
	}
	return false
}

/**
回文链表
*/
func isPalindrome(head *ListNode) bool {
	cur := head
	length := 0
	arr := make([]int, length)
	for cur != nil {
		arr = append(arr, cur.Val)
		length++
		cur = cur.Next
	}

	left := 0
	right := length - 1

	for left < right {
		if arr[left] == arr[right] {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}

/**
合并两个有序链表
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cur1 := list1
	cur2 := list2

	var newFront *ListNode
	var newRear *ListNode

	if cur1 == nil {
		return cur2
	}
	if cur2 == nil {
		return cur1
	}

	for cur1 != nil && cur2 != nil {
		if newFront == nil && newRear == nil {
			if cur1.Val <= cur2.Val {
				newFront = cur1
				newRear = cur1
				cur1 = cur1.Next
			} else {
				newFront = cur2
				newRear = cur2
				cur2 = cur2.Next
			}
		} else {
			if newRear != nil && cur1.Val <= cur2.Val {
				newRear.Next = cur1
				newRear = cur1
				cur1 = cur1.Next
			} else if newRear != nil && cur1.Val > cur2.Val {
				newRear.Next = cur2
				newRear = cur2
				cur2 = cur2.Next
			}
		}
	}

	if cur1 != nil && newRear != nil {
		newRear.Next = cur1
	}
	if cur2 != nil && newRear != nil {
		newRear.Next = cur2
	}
	return newFront
}

/**
反转链表
*/
func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	cur := head

	for cur != nil {
		curNext := cur.Next
		if newHead == nil {
			cur.Next = nil
			newHead = cur
		} else {
			cur.Next = newHead
			newHead = cur
		}
		cur = curNext
	}
	return newHead
}

// 1 -> 2 -> 3 -> 4 -> 5
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}

	preNode := length - n
	if preNode == 0 {
		return head.Next
	} else {
		cur := head
		i := 1
		for i < preNode {
			i++
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
	}
	return head
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	list := NewLinkList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	//list.Rear.Next = list.Front

	front := list.Front
	cycle := hasCycle2(front)
	fmt.Println(cycle)
}
