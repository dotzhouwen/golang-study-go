package main

import (
	"fmt"
	"time"
)

type Request interface {}
func handler(r Request) {
	fmt.Println(r.(int))
}

const RateLimitPeriod = time.Minute
const RateLimit = 200

func main() {

}

func handleRequests(request <- chan Request) {
	quotas := make(chan time.Time, RateLimit)

	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		defer tick.Stop()

		for t := range tick.C {
			select {
			case quotas <- t:
			default:

			}
		}
	}()
}
