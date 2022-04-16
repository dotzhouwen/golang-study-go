package main

import "fmt"

type Heap []int

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) siftUp(i int) {
	for {
		f := (i - 1) / 2
		if i == f || h.less(f, i) {
			break
		}
		h.swap(f, i)
		i = f
	}
}

func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.siftUp(len(*h) - 1)
}

func (h Heap) siftDown(i int) {
	for {
		l := 2*i + 1
		if l >= len(h) {
			break
		}

		j := l
		if r := l + 1; r < len(h) && h.less(r, l) {
			j = r
		}
		if h.less(i, j) {
			break
		}

		h.swap(i, j)
		i = j
	}
}

func (h *Heap) Pop() int {
	n := len(*h) - 1
	h.swap(0, n)
	x := (*h)[n]
	*h = (*h)[0:n]
	h.siftDown(0)
	return x
}

func (h Heap) Init() {
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}

func main() {
	h := Heap{20, 7, 3, 10, 15, 25, 30, 17, 19}
	h.Init()
	fmt.Println(h)

	for len(h) > 0 {
		pop := h.Pop()
		fmt.Println(pop)
	}
}
