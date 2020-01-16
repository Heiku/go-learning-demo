package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

type SpinLock struct {
	f uint32
}

func (s1 *SpinLock) Lock() {
	for !s1.TryLock() {
		runtime.Gosched()
	}
}

func (s1 *SpinLock) UnLock() {
	atomic.StoreUint32(&s1.f, 0)
}

func (sl *SpinLock) TryLock() bool {
	return atomic.CompareAndSwapUint32(&sl.f, 0, 1)
}

func (sl *SpinLock) String() string {
	if atomic.LoadUint32(&sl.f) == 1 {
		return "Locked"
	}
	return "Unlocked"
}

func main() {
	var mu = &SpinLock{}
	go func() {
		mu.Lock()
		time.Sleep(time.Second)
		mu.UnLock()
	}()

	time.Sleep(time.Second)

	ok := mu.TryLock()
	if ok {
		fmt.Println("got the lock")
		mu.UnLock()
		return
	} else {
		fmt.Println("can't get the lock")
	}
}
