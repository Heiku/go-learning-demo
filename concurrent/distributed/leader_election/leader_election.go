package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	endpoints := []string{"http://127.0.0.1:2379"}

	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	var electName = "/my-electName/"

	// watch electing process
	go watch(cli, electName)

	// 10 goroutine start electing
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go elect(i, cli, electName, &wg)
	}

	// find out who is the leader
	go query(cli, electName)
	wg.Wait()
}

func elect(id int, cli *clientv3.Client, electName string, wg *sync.WaitGroup) {
	defer wg.Done()

	// create a session
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()

	// create a election
	e1 := concurrency.NewElection(s1, electName)

	time.Sleep(time.Duration(5 * time.Second))

	// start electing
	log.Println("acampaigning for ID: ", id)
	if err := e1.Campaign(context.Background(), strconv.Itoa(id)); err != nil {
		log.Fatal(err)
	}
	log.Println("campaigned for ID: ", id)

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	// give up
	if err := e1.Resign(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Println("resigned for ID: ", id)
}

// observer
func watch(cli *clientv3.Client, elecName string) {
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()

	e1 := concurrency.NewElection(s1, elecName)
	ch := e1.Observe(context.TODO())

	for i := 0; i < 10; i++ {
		resp := <-ch
		log.Println("leader changed: to", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
	}
}

func query(cli *clientv3.Client, electName string) {
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()

	e1 := concurrency.NewElection(s1, electName)
	for i := 0; i < 10; i++ {
		resp, err := e1.Leader(context.Background())
		if err != nil {
			log.Printf("failed to get the current leader: %v", err)
			time.Sleep(9 * time.Second)
			continue
		}
		log.Println("current leader:", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
		time.Sleep(9 * time.Second)
	}
}
