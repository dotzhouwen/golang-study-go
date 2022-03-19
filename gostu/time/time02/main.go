package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("now:%v\n", now)

	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()

	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)

	timeObj := time.Unix(timestamp1, 0)
	fmt.Printf("timeObj:%v\n", timeObj)
}
