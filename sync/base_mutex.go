package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Counter struct {
	m sync.Mutex
	u uint64
}

func (c *Counter) Value() uint64 {
	c.m.Lock()
	defer c.m.Unlock()

	return c.u
}

func (c *Counter) Increase(delta uint64) {
	c.m.Lock()
	c.u++
	c.m.Unlock()
}

func main() {
	var c Counter
	for i := 0; i < 1000; i++ {
		go func() {
			for k := 0; k < 1000; k++ {
				c.Increase(1)
				fmt.Println(c.Value())
			}
		}()
	}

	for c.Value() < 10000000 {
		runtime.Gosched()
	}

}
