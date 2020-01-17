package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	endpoints := []string{"http://127.0.0.1:2379"}
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	var lockName = "my-lock"

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go startSession(i, cli, lockName, &wg)
	}

	wg.Wait()
}

func startSession(i int, cli *clientv3.Client, lockName string, wg *sync.WaitGroup) {
	defer wg.Done()

	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()

	m1 := concurrency.NewMutex(s1, lockName)

	// acquire lock
	log.Println("acquiring lock for ID: ", i)
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Println("acquired lock for ID: ", i)

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	if err := m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Println("release lock for ID: ", i)
}
