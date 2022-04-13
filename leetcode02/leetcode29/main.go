package main

import "fmt"

/**
设计哈希映射
*/
type MyHashMap struct {
	defaultSize int
	pairs       []*Pair
}

type Pair struct {
	Key   int
	Value int
	Next  *Pair
}

func Constructor() MyHashMap {
	return MyHashMap{
		defaultSize: 100,
		pairs:       make([]*Pair, 100),
	}
}

func (this *MyHashMap) index(key int) int {
	return key % this.defaultSize
}

func (this *MyHashMap) Put(key int, value int) {
	index := this.index(key)
	pair := this.pairs[index]
	if pair == nil {
		p := Pair{Key: key, Value: value}
		this.pairs[index] = &p
	} else {
		var prePair *Pair
		cur := pair
		for cur != nil {
			if cur.Key == key {
				cur.Value = value
				return
			}
			prePair = cur
			cur = cur.Next
		}
		if prePair != nil {
			newPair := Pair{Key: key, Value: value}
			prePair.Next = &newPair
		}
	}
}

func (this *MyHashMap) Get(key int) int {
	index := this.index(key)
	pair := this.pairs[index]
	cur := pair
	for cur != nil {
		if cur.Key == key {
			return cur.Value
		}
		cur = cur.Next
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	index := this.index(key)
	pair := this.pairs[index]

	if pair == nil {
		return
	} else {
		var prePair *Pair
		cur := pair
		for cur != nil {
			if cur.Key == key {
				if prePair == nil {
					this.pairs[index] = cur.Next
				} else {
					prePair.Next = cur.Next
				}
			}
			prePair = cur
			cur = cur.Next
		}
	}
}

func main() {
	hashMap := Constructor()
	hashMap.Put(1, 10)
	fmt.Println(hashMap.Get(1))

	hashMap.Put(1, 100)
	fmt.Println(hashMap.Get(1))

	hashMap.Put(101, 10001)
	fmt.Println(hashMap.Get(101))

	hashMap.Remove(1)
	fmt.Println(hashMap.Get(1))
	fmt.Println(hashMap.Get(101))

	hashMap.Remove(101)
	fmt.Println(hashMap.Get(101))
}
