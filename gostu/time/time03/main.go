package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	later := now.Add(time.Hour)
	fmt.Printf("%v\n", later)
}
