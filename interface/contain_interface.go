package main

import "fmt"

func main() {
	var i interface{}
	i = []int{1, 2, 4}
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	i = map[string]int{"Go": 2009}
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	i = true
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	i = 1
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	i = "abc"
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	i = nil
	fmt.Println(i)
	fmt.Printf("%T\n", i)
}
