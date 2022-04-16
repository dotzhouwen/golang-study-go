package main

import "fmt"

/**
两数之和 III - 数据结构设计
*/
type TwoSum struct {
	index int
	nums  []int
	hash  map[int]int
}

func Constructor() TwoSum {
	return TwoSum{
		index: 0,
		nums:  make([]int, 0),
		hash:  make(map[int]int, 0),
	}
}

func (this *TwoSum) Add(number int) {
	this.nums = append(this.nums, number)
	this.hash[number] = this.index
	this.index++
}

func (this *TwoSum) Find(value int) bool {
	for i, v := range this.nums {
		if ix, exists := this.hash[value-v]; exists {
			if i != ix {
				return true
			}
		}
	}
	return false
}

func main() {
	twoSum := Constructor()
	twoSum.Add(1)
	twoSum.Add(3)
	twoSum.Add(5)
	find := twoSum.Find(4)
	fmt.Println(find)

	fmt.Println(twoSum.Find(8))
	fmt.Println(twoSum.Find(9))
}
