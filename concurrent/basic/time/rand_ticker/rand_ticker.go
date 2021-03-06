package main

import (
	"log"
	"math/rand"
	"time"
)

type RandTicker struct {
	C    <-chan time.Time
	done chan struct{}
}

func NewRandTicker(d, variance time.Duration) *RandTicker {
	c := make(chan time.Time, 1)
	done := make(chan struct{})

	go start(d, variance, c, done)

	return &RandTicker{
		C:    c,
		done: done,
	}
}

func start(d, variance time.Duration, c chan time.Time, done chan struct{}) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		vr := time.Duration(rnd.Int63n(int64(2*variance)) - int64(variance))
		tmr := time.NewTimer(d + vr)
		select {
		case fireAt := <-tmr.C:
			select {
			case c <- fireAt:
			default:
			}
		case <-done:
			tmr.Stop()
			return
		}
	}
}

func (rtk *RandTicker) Stop() {
	close(rtk.done)
}

func main() {
	time.Now().String()

	t := NewRandTicker(10*time.Second, 2*time.Second)
	defer t.Stop()

	for i := 0; i < 10; i++ {
		fired := <-t.C
		log.Println(fired.Format("fired at 12:03:36"))
	}
}
