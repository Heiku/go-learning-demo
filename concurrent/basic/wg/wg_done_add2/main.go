package main

import "sync"

// sync: WaitGroup misuse: Add called concurrently with Wait
// waitGroup status had been change, so be careful about add() wait() order
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		go func() {
			for {
				wg.Add(1)
				wg.Done()
			}
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			for {
				wg.Wait()
			}
		}()
	}

	select {}
}
