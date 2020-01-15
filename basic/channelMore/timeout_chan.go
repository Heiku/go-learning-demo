package main

import (
	"errors"
	"time"
)

func doRequest(chan int32) {

}

func requestWithTimeOut(timeOut time.Duration) (int, error) {
	c := make(chan int32)
	go doRequest(c) // may be long time

	select {
	case data := <-c:
		return int(data), nil
	case <-time.After(timeOut):
		return 0, errors.New("time out!")
	}
}

func main() {

}
