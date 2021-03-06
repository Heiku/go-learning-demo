package main

import (
	"context"
	"fmt"
	"time"
)

// context with value,
var key string = "name"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	valueCtx := context.WithValue(ctx, key, "[first]")
	go watch(valueCtx)

	time.Sleep(10 * time.Second)
	cancel()

}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), " watching exit, stopping...")
			return
		default:
			fmt.Println(ctx.Value(key), " watching now...")
			time.Sleep(2 * time.Second)
		}
	}
}
