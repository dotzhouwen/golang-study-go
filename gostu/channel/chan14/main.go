package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var rnd int32

	select {
	case rnd = <-source():
	case rnd = <-source():
	case rnd = <-source():
	}

	fmt.Println(rnd)
}

func source() <-chan int32 {
	c := make(chan int32, 1)
	go func() {
		ra, rb := rand.Int31(), rand.Intn(3) + 1
		time.Sleep(time.Duration(rb) * time.Second)
		c <- ra
	}()
	return c
}
