package main

import (
	"context"
	"fmt"
	"time"
)

// use context watch multi-goroutine
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "[first]")
	go watch(ctx, "[second]")
	go watch(ctx, "[third]")

	time.Sleep(10 * time.Second)
	fmt.Println("watching stopped")
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " watching exit, stopping...")
			return
		default:
			fmt.Println(name, " watching now...")
			time.Sleep(2 * time.Second)
		}
	}
}
