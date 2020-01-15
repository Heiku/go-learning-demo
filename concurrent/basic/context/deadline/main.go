package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx1, c1 := context.WithCancel(context.Background())
	go func() {
		fmt.Println("g1 start")
		<-ctx1.Done()
		fmt.Println("g1 done,err: ", ctx1.Err())
	}()

	// deadline did't work, because  parent ctx exit
	// parent exit, child context must exit
	ctx2, c2 := context.WithDeadline(ctx1, time.Now().Add(time.Second*10))
	go func() {
		fmt.Println("g2 start")
		<-ctx2.Done()
		fmt.Println("g2 done,err: ", ctx2.Err())
	}()

	time.Sleep(time.Second)
	c1()

	time.Sleep(10 * time.Second)
	c2()
}
