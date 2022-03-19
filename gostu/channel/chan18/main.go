package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	c := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			tmp := rand.Intn(10)
			time.Sleep(time.Duration(tmp) * time.Second)
			fmt.Println("I want to sleep", tmp, "second")
			c <- i
		}(i)
	}

	for {
		select {
		case i := <-c:
			fmt.Printf("The %d goroutine is done.\n", i)
		case <-ticker.C:
			fmt.Println("Time to go out")
			os.Exit(5)
		}
	}
}
