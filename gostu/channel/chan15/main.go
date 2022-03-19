package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	for range Tick(time.Second) {
		fmt.Println(time.Since(t))
	}
}

func Tick(d time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		for {
			time.Sleep(d)
			select {
			case c <- struct{}{}:
			default:
			}
		}
	}()
	return c
}
