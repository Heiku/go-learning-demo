package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("exit correctly!")
	}()

	fmt.Println("hi!")

	defer func() {
		v := recover()  // recover() receive panic info and return
		fmt.Println("panic recover(): ", v)
	}()

	panic("bye bye !")
	fmt.Println("can not reached")
}
