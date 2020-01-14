package main

import (
	"sync"
	"time"
)

var mu sync.Mutex

func lock() {
	mu.Lock()
	defer mu.Lock()
}

func main() {
	go lock()
	time.Sleep(1e9)
}
