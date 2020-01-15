package main

import (
	"fmt"
	"time"
)

func main() {
	var ball = make(chan string)

	kickBall := func(playerName string) {
		for {
			fmt.Print(<-ball, "pass ball", "\n")
			time.Sleep(time.Second)

			ball <- playerName
		}
	}

	go kickBall("batman")
	go kickBall("spider man")
	go kickBall("super man")
	go kickBall("iron man")

	ball <- "start"

	var c chan bool
	<-c // always block
}
