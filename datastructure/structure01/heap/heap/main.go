package main

import "fmt"

func adjustHeap(a []int, pos int, length int) {
	for {
		child := pos*2 + 1
		if child >= length-1 {
			break
		}

		if a[child+1] > a[child] {
			child += 1
		}

		if a[pos] < a[child] {
			a[pos], a[child] = a[child], a[pos]
			pos = child
		} else {
			break
		}
	}
}

func buildHeap(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		adjustHeap(a, i, len(a))
	}
}

func heapSort(a []int) {
	buildHeap(a)
	fmt.Println("build heap over:", a)
	fmt.Println("===================")

	for i := len(a) - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		adjustHeap(a, 0, i)
	}
}

func main() {
	num := []int{98, 48, 777, 63, 57, 433, 23, 1112, 1}
	heapSort(num)
	fmt.Println(num)
}
