package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctx1, c1 := context.WithCancel(ctx)
	go func() {
		fmt.Println("g1 start")
		<-ctx1.Done()
		fmt.Println("g1 done, err:", ctx1.Err())
	}()

	ctx2 := context.WithValue(ctx1, "aaa", 1)
	go func() {
		fmt.Println("g2 start")
		<-ctx2.Done()
		fmt.Println("g2 done, err:", ctx1.Err())
	}()

	ctx3, c3 := context.WithCancel(ctx2)
	go func() {
		fmt.Println("g3 start")
		<-ctx3.Done()
		fmt.Println("g3 done, err:", ctx1.Err())
	}()

	time.Sleep(1e9)
	c1()
	time.Sleep(5 * time.Second)
	c3()
}
