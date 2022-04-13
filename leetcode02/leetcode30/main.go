package main

import "fmt"

/**
只出现一次的数字
*/
func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res = res ^ v
	}
	return res
}

func main() {
	a := []int{2, 3, 2, 4, 4}
	fmt.Println(singleNumber(a))
}
