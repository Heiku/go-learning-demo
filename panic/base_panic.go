package main

import "fmt"

// when method produce a panic, we can consider the method invoke runtime.Goexit()
// panic can be recover, but goexit() can't be cancelled
func main() {
	defer func() {
		fmt.Println("exit correctly!")
	}()

	fmt.Println("hi!")

	defer func() {
		v := recover() // recover() receive panic info and return
		fmt.Println("panic recover(): ", v)
	}()

	panic("bye bye !")
	fmt.Println("can not reached")
}
