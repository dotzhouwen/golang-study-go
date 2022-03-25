package main

import "fmt"

func mergeSort(nums []int) {

}

func merge(left []int, right []int) {

	var result []int

	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	fmt.Printf("i:%d\n", i)
	fmt.Printf("j:%d\n", j)

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

}

func main() {
	left := []int{2, 4, 6}
	right := []int{1, 3, 5}
	var result []int

	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	fmt.Printf("i:%d\n", i)
	fmt.Printf("j:%d\n", j)

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	fmt.Println(result)
}
