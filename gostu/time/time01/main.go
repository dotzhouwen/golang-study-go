package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)

	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())

	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
}
