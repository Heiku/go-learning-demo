package main

import (
	"fmt"
	"time"
)

type Request interface {
}

func handle(t time.Time) {
	fmt.Println(t)
}

const RateLimitPeriod = time.Minute
const RateLimit = 200

func handlerRequests(requests <-chan Request) {
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

	for _ = range requests {
		t := <-quotas
		go handle(t)
	}
}

func main() {
	requests := make(chan Request)
	go handlerRequests(requests)

	for i := 0; ; i++ {
		requests <- i
	}
}
