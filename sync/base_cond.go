package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const N = 10
	var values [N]string

	cond := sync.NewCond(&sync.Mutex{})
	cond.L.Lock()

	for i := 0; i < N; i++ {
		d := time.Second * time.Duration(rand.Intn(10)) / 10

		go func(i int) {
			time.Sleep(d)
			cond.L.Lock()
			values[i] = string('a' + 1)
			cond.Broadcast()
			cond.L.Unlock()
		}(i)
	}

	checkCondition := func() bool {
		fmt.Println(values)
		for i := 0; i < N; i++ {
			if values[i] == "" {
				return false
			}
		}
		return true
	}
	for !checkCondition() {
		cond.Wait()
	}
	cond.L.Unlock()
}
