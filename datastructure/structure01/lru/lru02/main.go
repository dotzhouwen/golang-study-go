package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Prev  *Node
	Next  *Node
	Key   string
	Value interface{}
}

type LRUCache struct {
	cache    map[string]*Node
	head     *Node
	tail     *Node
	len      uint64
	capacity uint64
}

func NewNode(key string, value interface{}) *Node {
	return &Node{
		Prev:  nil,
		Next:  nil,
		Key:   key,
		Value: value,
	}
}

func NewLRUCache(cap uint64) *LRUCache {
	return &LRUCache{
		cache:    make(map[string]*Node),
		head:     nil,
		tail:     nil,
		capacity: cap,
	}
}

func (c *LRUCache) AddHead(node *Node) {
	if c.len == 0 {
		c.head = node
		c.tail = node
	} else {
		_node := c.head
		node.Next = _node
		_node.Prev = node
		c.head = node
	}
	c.len++
}

func (c *LRUCache) Remove(node *Node) {
	if c.head == node {
		//c.head.Next = node.Next
		//node.Next.Prev = c.head
		c.head = c.head.Next
		c.head.Prev = nil
	} else if c.tail == node {
		c.tail = c.tail.Prev
		c.tail.Next = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
	delete(c.cache, node.Key)
	c.len--
}

func (c *LRUCache) Get(key string) interface{} {
	if node, ok := c.cache[key]; ok {
		c.Remove(node)
		c.AddHead(node)
		return node.Value
	}
	return nil
}

func (c *LRUCache) Put(key string, value interface{}) {
	if node, ok := c.cache[key]; ok {
		node.Value = value
		c.Remove(node)
		c.AddHead(node)
	} else {
		newNode := NewNode(key, value)
		if c.len == c.capacity {
			c.Remove(c.tail)
		}
		c.AddHead(newNode)
		c.cache[key] = newNode
	}
}

func (c *LRUCache) Print() {
	builder := strings.Builder{}
	cur := c.head
	for cur != nil {
		builder.WriteString(fmt.Sprintf("%v -> %v", cur.Key, cur.Value))
		cur = cur.Next
		if cur != nil {
			builder.WriteString(" , ")
		}
	}
	fmt.Println(builder.String())
}

func main() {
	cache := NewLRUCache(5)

	cache.Put("John", 20)
	cache.Put("Marry", 30)
	cache.Put("Jack", 25)

	cache.Print()

	ret := cache.Get("Marry")
	fmt.Println(ret)
	cache.Print()

	fmt.Println("-----")

	ret = cache.Get("John")
	fmt.Println(ret)
	cache.Print()

	cache.Put("Jack", 920)
	cache.Print()

	cache.Put("Stive", 100)
	cache.Print()

	cache.Put("dot", 100)
	cache.Print()

	cache.Put("JackMa", 300)
	cache.Print()
	//cache.AddHead(NewNode("john", 30))
	//cache.AddHead(NewNode("marry", 30))
	//cache.AddHead(NewNode("jack", 90))

	//cache.Print()
}
