package main

import (
	"fmt"
	recipe "github.com/coreos/etcd/contrib/recipes"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// Queue implements a multi-reader, multi-writer distributed queue.
func main() {
	rand.Seed(time.Now().UnixNano())

	endpoints := []string{"http://127.0.0.1:2379"}
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	var queueName = "my-queue"

	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go write(i, cli, queueName, &wg)
	}

	for i := 0; i < 10; i++ {
		go read(10+i, cli, queueName, &wg)
	}

	wg.Wait()
}

func write(id int, cli *clientv3.Client, queueName string, wg *sync.WaitGroup) {
	defer wg.Done()

	q := recipe.NewQueue(cli, queueName)

	for i := 0; i < 10; i++ {
		q.Enqueue(fmt.Sprintf("g-%d-key-%d", id, i))
	}
}

func read(id int, cli *clientv3.Client, queueName string, wg *sync.WaitGroup) {
	defer wg.Done()

	q := recipe.NewQueue(cli, queueName)

	for i := 0; i < 10; i++ {
		v, err := q.Dequeue()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("goroutine %d received: %s\n", id, v)
	}
}