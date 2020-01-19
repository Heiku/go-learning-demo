package main

import (
	"sync"
	"sync/atomic"
	"time"
)

var a uint64
var mu sync.Mutex

func foo() {
	if a == 1 {
		return
	}
	mu.Lock()
	if a == 0 {
		a = 1
	}
	mu.Unlock()

}

func fooByAtomic() {
	if atomic.LoadUint64(&a) == 1 {
		return
	}
	mu.Lock()
	if a == 0 {
		atomic.StoreUint64(&a, 1)
	}
	mu.Unlock()
}

func fooByMoreScope() {
	mu.Lock()
	defer mu.Unlock()

	if a == 1 {
		return
	}
	a = 1
}

func main() {
	go foo()
	go foo()

	// go fooByAtomic()
	// go fooByAtomic()

	// go fooByMoreScope()
	// go fooByMoreScope()

	time.Sleep(time.Second)
}
