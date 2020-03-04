package main

import "fmt"

func main() {
	fmt.Println(make(map[string]int))

	// make create map
	m := make(map[string]int, 3)
	fmt.Println(m, len(m))

	m["C"] = 1972
	m["Go"] = 2009
	m["Java"] = 0
	m["Erlang"] = 1
	// automatic expansion
	fmt.Println(m, len(m))

	// make create slice	// also	 can use new()
	// type, len, capacity
	s := make([]int, 3, 5)
	fmt.Println(s, len(s), cap(s))

	s = make([]int, 2)
	fmt.Println(s, len(s), cap(s))
}
