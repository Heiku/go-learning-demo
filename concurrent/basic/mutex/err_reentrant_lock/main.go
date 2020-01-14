package main

import (
	"fmt"
	"sync"
)

// Mutex is exclusive lock

func M1(locker sync.Locker) {
	fmt.Println("in M1")
	locker.Lock()
	M2(locker)
	locker.Unlock()
}

func M2(l sync.Locker) {
	fmt.Println("in M2")
	M3(l)
}

func M3(l sync.Locker) {
	l.Lock()
	fmt.Println("in M3")
	l.Unlock()
}

func main() {
	l := &sync.Mutex{}
	M1(l)
}
