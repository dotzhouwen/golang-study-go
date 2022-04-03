package main

import "fmt"

func test(m map[string]int) {
	m["john"] = 20
	m["alice"] = 30
	m["Jack"] = 32
	for i := 0; i < 100; i++ {
		m[fmt.Sprintf("test-%d", i)] = i
	}
}

func main() {
	m := make(map[string]int, 1)
	test(m)
	fmt.Println(m)
}
