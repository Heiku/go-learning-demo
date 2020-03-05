package main

import (
	"fmt"
	"time"
)

// ticker like a timer, when put timeData into ticker.C on schedule
func main() {
	ticker := time.NewTicker(time.Second * 1)
	/*	go func() {
			for t := range ticker.C {
				fmt.Print("Tick at", t, "\n")
			}
			fmt.Println()
		}()
		select {}*/

	// can use for-range or for select to traversal channel
	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Print("Tick at ", t, "\n")
			}
		}
	}()
	select {}
}
