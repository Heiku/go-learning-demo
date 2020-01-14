package error

import "sync/atomic"

// maybe wrong in multi-goroutine
type BadOnce struct {
	done uint32
}

func (o *Once) BadDo(f func()) {
	if !atomic.CompareAndSwapUint32(&o.done, 0, 1) {
		return
	}
	f()
}
