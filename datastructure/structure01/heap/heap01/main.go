package main

import (
	"fmt"
	"strings"
)

type Heap struct {
	Size  int
	Elems []int
}

func NewHeap(maxSize int) *Heap {
	minHead := Heap{Size: -1, Elems: make([]int, maxSize, maxSize)}
	return &minHead
}

func (h *Heap) Print() {
	builder := strings.Builder{}
	for i := 0; i <= h.Size; i++ {
		builder.WriteString(fmt.Sprintf("%d ", h.Elems[i]))
	}
	fmt.Println(builder.String())
}

func (h *Heap) Push(x int) {
	if h.Size == len(h.Elems)-1 {
		return
	}
	h.Size++
	h.Elems[h.Size] = x

	pos := h.Size
	for {
		if pos <= 0 {
			break
		}
		parent := (pos - 1) / 2
		if h.Elems[parent] <= x {
			break
		}
		h.Elems[parent], h.Elems[pos] = h.Elems[pos], h.Elems[parent]
		pos = parent
	}
}

func (h *Heap) IsEmpty() bool {
	return h.Size == -1
}

func (h *Heap) Pop() int {
	if h.Size == -1 {
		panic("The heap is null")
	}
	h.Elems[h.Size], h.Elems[0] = h.Elems[0], h.Elems[h.Size]
	v := h.Elems[h.Size]
	h.Size--

	pos := 0
	for {
		child := pos*2 + 1
		if child > h.Size {
			break
		}
		if child+1 <= h.Size && h.Elems[child+1] < h.Elems[child] {
			child = child + 1
		}
		if h.Elems[pos] > h.Elems[child] {
			h.Elems[pos], h.Elems[child] = h.Elems[child], h.Elems[pos]
			pos = child
		} else {
			break
		}
	}
	return v
}

func main() {
	h := NewHeap(10)
	h.Push(10)
	h.Push(9)
	h.Push(4)
	h.Push(8)
	h.Push(3)

	h.Print()

	for !h.IsEmpty() {
		pop := h.Pop()
		fmt.Println("pop:", pop)
	}
}
