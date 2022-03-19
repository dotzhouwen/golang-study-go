package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(c int) {
	log.Print("customer#", c, " enters the bar")

	seat := <-bar

	log.Print("++customer#", c, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(2 + rand.Intn(10)))
	log.Print("--customer#", c, " frees seat#", seat)

	bar <-seat
}

func main() {
	rand.Seed(time.Now().UnixNano())
	bar24x7 := make(Bar, 10)

	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		fmt.Println("create goroutine")
		go bar24x7.ServeCustomer(customerId)
	}

	for {
		time.Sleep(time.Second)
	}
}
