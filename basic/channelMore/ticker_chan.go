package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Print("Tick at", t, "  ")
		}
		fmt.Println()
	}()
	select {}
}
