package main

import (
	"context"
	"fmt"
	"github.com/marusama/cyclicbarrier"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	cnt := 0
	barrier := cyclicbarrier.NewWithAction(10, func() error {
		cnt++
		return nil
	})

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			defer wg.Done()

			for j := 0; j < 5; j++ {
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				log.Printf("goroutine %d waits", i)

				err := barrier.Await(context.TODO())
				log.Printf("goroutine %d is ok", i)

				if err != nil {
					panic(err)
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}
