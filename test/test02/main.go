package main

import "fmt"

func test(arr *[]int) {
	*arr = append(*arr, 1)
}

func main() {
	arr := make([]int, 0)
	test(&arr)
	fmt.Println(arr)
}
