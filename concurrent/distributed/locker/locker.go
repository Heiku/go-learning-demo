package main

import (
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

func startSession(id int, cli *clientv3.Client, lockName string, wg *sync.WaitGroup) {
	defer wg.Done()

	// new session
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()

	locker := concurrency.NewLocker(s1, lockName)

	// start acquiring lock
	log.Println("acquiring lock for ID: ", id)
	locker.Lock()
	log.Println("acquired lock for ID: ", id)

	// release lock
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	locker.Unlock()
	log.Println("released lock for ID: ", id)
}
