package error

import (
	"sync"
	"sync/atomic"
)

// if variable 'done' set 1, it means the func had been called
// atomic can be used to cas、volatile
type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.slowDo(f)
}

func (o *Once) slowDo(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()

	var err error
	if o.done == 0 {
		err = f()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}
