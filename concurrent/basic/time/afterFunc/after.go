package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("fired")
	})

	// When the Timer expires, the current time will be sent on C except AfterFunc
	// so chan is empty, t = nil
	t := <-timer.C
	log.Printf("fired at %s", t.String())
}
