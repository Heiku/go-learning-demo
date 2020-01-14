package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// simulate the mutex process
const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type Mutex struct {
	sync.Mutex
}

func (m *Mutex) TryLock() bool {
	// unsafe.Pointer return a pointer  -> sync.Mutex
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}

	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}

func (m *Mutex) Count() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v >> mutexWaiterShift
	v = v + (v & mutexLocked)
	return int(v)
}

func (m *Mutex) IsWoken() bool {
	start := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return start&mutexWoken == mutexWoken
}

func (m *Mutex) IsStarving() bool {
	start := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return start&mutexStarving == mutexStarving
}

// test
func try() {
	var mu Mutex
	go func() {
		mu.Lock()
		time.Sleep(time.Second)
		mu.Unlock()
	}()
	time.Sleep(time.Second)

	ok := mu.TryLock()
	if ok {
		fmt.Println("got the lock")
		mu.Unlock()
		return
	}
	fmt.Println("can't get the lock")
}

func count() {
	var mu Mutex
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Lock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("waiting:%d, woken: %t, starving: %t\n", mu.Count(), mu.IsWoken(), mu.IsStarving())
}

func main() {
	fmt.Println(mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift)
	try()
	count()
}
