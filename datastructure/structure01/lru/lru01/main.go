package main

import "fmt"

type Node struct {
	Prev  *Node
	Next  *Node
	key   int
	value int
}

func NewNode(k, v int) *Node {
	return &Node{
		key:   k,
		value: v,
	}
}

type LRUCache struct {
	cache map[int]*Node
	head  *Node
	tail  *Node
	len   int
	cap   int
}

func (c *LRUCache) Print() {
	str := ""
	p := c.head
	for p != nil {
		str += fmt.Sprintf("key:%d, value:%d", p.key, p.value)
		p = p.Next
		if p != nil {
			str += fmt.Sprintf(" => ")
		}
	}
	fmt.Println("双链表:" + str)
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		cache: make(map[int]*Node),
		cap:   capacity,
	}
}

func (c *LRUCache) AddHead(node *Node) {
	if c.len == 0 {
		c.head = node
		c.tail = node
	} else {
		_node := c.head
		c.head = node
		_node.Prev = node
		c.head.Next = _node
	}
	c.len++
}

func (c *LRUCache) Remove(n *Node) {
	if c.head == n {
		c.head = c.head.Next
		c.head.Prev = nil
	} else if c.tail == n {
		c.tail = c.tail.Prev
		c.tail.Next = nil
	} else {
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev
	}
	delete(c.cache, n.key)
	c.len--
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.Remove(node)
		c.AddHead(node)
		return node.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.cache[key]; ok {
		c.Get(key)
	} else {
		if c.len == c.cap {
			c.Remove(c.tail)
		}
		node := NewNode(key, value)
		c.AddHead(node)
		c.cache[key] = node
	}
}

func main() {
	var result int
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Print()

	result = cache.Get(1)
	fmt.Println("result:", result)
	cache.Print()

	cache.Put(3, 3)
	cache.Print()

	result = cache.Get(2)
	fmt.Println(result)

	cache.Put(4, 4)
	cache.Print()

	cache.Get(3)
	cache.Print()
}
