package main

import (
	"fmt"
	"strings"
)

func main() {
	var helloworld = "hello-world"

	var hello = helloworld[:5]
	fmt.Println(hello[0])         // h unicode code = 104
	fmt.Printf("%T \n", hello[0]) // uint8

	// can't compile (it mean underlay array can not be changed)
	// hello[0] = 'H'
	// fmt.Println(&hello[0])

	fmt.Println(len(hello), len(helloworld), strings.HasPrefix(helloworld, hello))

	// append string
	hi := []byte("Hello ")
	world := "world !"

	// helloWorld := append(hi, []byte(world)...)
	helloWorld := append(hi, world...)
	fmt.Println(string(helloWorld))

}
