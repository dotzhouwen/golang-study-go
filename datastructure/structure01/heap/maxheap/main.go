package main

import (
	"fmt"
	"strings"
)

type MaxHeap struct {
	Data     []int
	Capacity int
	Size     int
}

func NewMaxHeap(capacity int) *MaxHeap {
	return &MaxHeap{
		Data:     make([]int, capacity, capacity),
		Capacity: capacity,
		Size:     0,
	}
}

func (h *MaxHeap) Print() {
	sb := strings.Builder{}
	for i := 0; i < h.Size; i++ {
		if i != h.Size-1 {
			sb.WriteString(fmt.Sprintf("%d, ", h.Data[i]))
		} else {
			sb.WriteString(fmt.Sprintf("%d", h.Data[i]))
		}
	}
	fmt.Println(sb.String())
}

func (h *MaxHeap) Add(val int) {
	h.Data[h.Size] = val
	h.Size += 1
	h.siftUp(h.Size - 1)
}

func (h *MaxHeap) FindMax() int {
	return h.Data[0]
}

func (h *MaxHeap) ExtractMax() int {
	first := h.Data[0]
	h.Data[0], h.Data[h.Size-1] = h.Data[h.Size-1], h.Data[0]
	h.Size -= 1
	h.siftDown(0)
	return first
}

func (h *MaxHeap) siftDown(index int) {
	left := h.leftChild(index)
	right := h.rightChild(index)

	for left < h.Size {
		child := left
		if right < h.Size && h.Data[right] > h.Data[left] {
			child += 1
		}
		if h.Data[index] > h.Data[child] {
			break
		}
		h.Data[child], h.Data[index] = h.Data[index], h.Data[child]
		index = child
		left = h.leftChild(index)
		right = h.rightChild(index)
	}
}

func (h *MaxHeap) siftUp(index int) {
	if index >= 0 {
		for index > 0 && h.Data[index] > h.Data[h.parent(index)] {
			h.Data[h.parent(index)], h.Data[index] = h.Data[index], h.Data[h.parent(index)]
			index = h.parent(index)
		}
	}
}

func (h *MaxHeap) leftChild(index int) int {
	return (index * 2) + 1
}

func (h *MaxHeap) rightChild(index int) int {
	return (index * 2) + 2
}

func (h *MaxHeap) parent(index int) int {
	if index < 1 {
		panic("index must be more than one")
	}
	return (index - 1) / 2
}

func main() {
	heap := NewMaxHeap(10)
	heap.Add(1)
	heap.Add(2)
	heap.Add(10)
	heap.Add(8)
	heap.Add(30)
	heap.Add(120)
	heap.Add(90)
	heap.Add(389)

	heap.Print()
	max := heap.FindMax()
	fmt.Println("max:", max)

	extractMax := heap.ExtractMax()
	fmt.Println("extractMax:", extractMax)
	heap.Print()

	extractMax = heap.ExtractMax()
	fmt.Println("extract:", extractMax)
	heap.Print()

	extractMax = heap.ExtractMax()
	fmt.Println("extract:", extractMax)
	heap.Print()

	extractMax = heap.ExtractMax()
	fmt.Println("extract:", extractMax)
	heap.Print()

	extractMax = heap.ExtractMax()
	fmt.Println("extract:", extractMax)
	heap.Print()

	extractMax = heap.ExtractMax()
	fmt.Println("extract:", extractMax)
	heap.Print()
}
