package main

import "fmt"

func main() {
	var ans []int
	fmt.Println(len(ans))

	for i := 0; i < 10; i++ {
		ans = append(ans, i)
		fmt.Println("len:", len(ans))
	}
}
