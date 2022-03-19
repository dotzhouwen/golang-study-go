package main

import "fmt"

type CNode struct {
	data interface{}
	next *CNode
}

type CList struct {
	size uint64
	head *CNode
}

func (c *CList) Print() {
	format := ""
	cur := c.GetHead()

	for ; cur.next != c.head; cur = cur.next {
		format += fmt.Sprintf("%v => ", cur.data)
	}

	format += fmt.Sprintf("%v", cur.data)
	fmt.Println(format)
}

func (c *CList) Init() {
	c.size = 0
	c.head = nil
}

func (c *CList) GetSize() uint64 {
	return c.size
}

func (c *CList) GetHead() *CNode {
	return c.head
}

func (c *CList) Append(data interface{}) bool {
	node := new(CNode)
	node.data = data

	if c.GetSize() == 0 {
		c.head = node
	} else {
		cur := c.GetHead()
		for ; cur.next != c.GetHead(); cur = cur.next {}
		cur.next = node
	}
	node.next = c.GetHead()
	c.size++
	return true
}

func (c *CList) InsertNext(elem *CNode, data interface{}) bool {
	if elem == nil {
		return false
	}
	node := new(CNode)
	node.data = data

	node.next = elem.next
	elem.next = node

	c.size++
	return true
}

func (c *CList) Remove(elem *CNode) interface{} {
	if elem == nil {
		return nil
	}
	cur := c.GetHead()
	for ; cur.next != elem ; cur = cur.next {}

	cur.next = elem.next
	c.size--
	return elem.data
}



func main() {
	list := new(CList)

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	list.Print()

	fmt.Println(list.GetSize())
}
