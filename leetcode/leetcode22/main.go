package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	result := 0
	for math.Abs(float64(x)) > 0 {
		mod := x % 10
		result = result*10 + mod
		x = x / 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

func main() {

	result := reverse(-123)
	fmt.Println(result)
	//fmt.Println(1234 % 10)
	//fmt.Println(1234 / 10 % 10)
	//fmt.Println(1234 / 10 /10 % 10)
	//fmt.Println(1234 / 10 / 10 / 10 % 10)
	//fmt.Println(1234 / 10 / 10 / 10 / 10 % 10)
	//fmt.Println(1234 / 10 / 10 / 10 / 10 / 10 % 10)
}
