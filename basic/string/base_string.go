package main

import "fmt"

/**
type _string struct{
	element *byte	// byte
	len				// int
}
*/
func main() {
	const world = "world"
	var hello = "hello"

	// append string
	var helloworld = hello + " " + world
	helloworld += "!"
	fmt.Println(helloworld)

	fmt.Println(hello == "hello")
	fmt.Println(hello > helloworld)
}
