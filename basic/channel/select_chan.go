package main

import "fmt"

// select case in nil channel will be blocked,
func main() {
	c := make(chan string, 2)

	trySend := func(v string) {
		select {
		case c <- v:
			fmt.Println("select send: ", v)
		default:
			// if buffer queue is full, will do this
			fmt.Println("send failed, buffer queue is full")
		}
	}

	tryReceive := func() string {
		select {
		case v := <-c:
			return v
		default:
			return "-" // if buffer queue is empty, will do this
		}
	}

	trySend("Hello")
	trySend("Hi")
	trySend("Byte!") // send failed

	fmt.Println(tryReceive())
	fmt.Println(tryReceive())

	// receive failed
	fmt.Println(tryReceive())
}
