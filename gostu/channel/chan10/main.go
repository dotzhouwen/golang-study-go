package main

import (
	"log"
	"math/rand"
	"time"
)

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomerAtSeat(c int, seat Seat) {
	log.Print("custome#", c, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Print("<-customer#", c, " frees seat#", seat)
	bar <- seat
}

func main() {
	rand.Seed(time.Now().UnixNano())
	bar24x7 := make(Bar, 10)
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		seat := <-bar24x7
		go bar24x7.ServeCustomerAtSeat(customerId, seat)
	}

	for {
		time.Sleep(time.Second)
	}
}
