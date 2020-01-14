package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once
	var count = 0

	go func() {
		defer func() {
			count++
			recover()
		}()

		once.Do(func() {
			fmt.Println("exec Do")
			count = 1 / count
		})

		fmt.Println("did't reached here")
	}()

	time.Sleep(time.Second)

	// once be used, did't called this method
	once.Do(func() {
		fmt.Println("exec here")
		count = 1 / count
	})

	// goroutine exit, didn't influence main thread
	fmt.Println(count)
	fmt.Println("end")
}
