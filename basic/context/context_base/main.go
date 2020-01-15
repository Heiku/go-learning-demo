package main

import (
	"context"
	"fmt"
	"time"
)

// context.Background() -> empty tree
// then, context.withCancel(parent), new a child tree (child context)
func main() {
	// child context, cancelFunc
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("watching exit, stopping...")
				return
			default:
				fmt.Println("watching...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("fine, watching stopping...")
	cancel()

	time.Sleep(5 * time.Second)
}
