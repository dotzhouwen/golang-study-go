package main

import (
	"log"
	"math/rand"
	"time"
)

type Customer struct {
	id int
}

type Bar chan Customer

func (bar Bar) ServeCustomer(c Customer) {
	log.Print("++ customer#", c.id, " starts drinks")
	time.Sleep(time.Second * time.Duration(3+rand.Intn(16)))
	log.Print("-- customer#", c.id, " leaves the bar")
	<-bar
}

func main() {
	rand.Seed(time.Now().UnixNano())
	bar24x7 := make(Bar, 10)

	for customerId := 0; ; customerId++ {
		//time.Sleep(time.Second * 2)
		customer := Customer{customerId}

		bar24x7 <- customer
		go bar24x7.ServeCustomer(customer)
	}

	for {
		time.Sleep(time.Second)
	}
}
