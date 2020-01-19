package main

import (
	"fmt"
	"time"
)

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := &Mutex{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

func (m *Mutex) Lock() {
	<-m.ch
}

func (m *Mutex) UnLock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

func (m *Mutex) TryLock(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true

		// after timeout, time.after channel will get a timestamp, when it output the ts, break the select
	case <-time.After(timeout):
	}
	return false
}

func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 0
}

func main() {
	m := NewMutex()
	ok := m.TryLock(time.Second)
	fmt.Printf("locked %v\n", ok)
	ok = m.TryLock(time.Second)
	fmt.Printf("locked %v\n", ok)
}
