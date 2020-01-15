package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go ClientHandler(con)
	}
}

func ClientHandler(c net.Conn) {
	defer func() {
		if v := recover(); v != nil {
			log.Println("catch a panic: ", v)
			log.Println("stop server stopping")
		}
		c.Close()
	}()
	panic("unknown error")
}
