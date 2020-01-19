package main

import "fmt"

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := Mutex{ch: make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return &mu
}

func (m *Mutex) Lock() {
	// when channel is empty(lock status), the called goroutine will be blocked
	<-m.ch
}

func (m *Mutex) UnLock() {
	select {
	// channel is full, it means the mutex is unlock status
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

func (m *Mutex) IsLock() bool {
	return len(m.ch) == 0
}

func main() {
	m := NewMutex()
	ok := m.TryLock()
	fmt.Printf("locked %v\n", ok)

	ok = m.TryLock()
	fmt.Printf("locked %v\n", ok)
}
