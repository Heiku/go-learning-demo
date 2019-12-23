package main

import "fmt"

func main() {
	m := map[string]int{"C": 1972, "C++": 1983, "Go": 2009}
	for s, i := range m {
		fmt.Printf("%v : %v\n", s, i)
	}

	a := [...]int{2, 3, 5, 7, 11}
	for i, v := range a {
		fmt.Printf("%v : %v\n", i, v)
	}

	s := []string{"go", "defer", "var", "goto"}
	for i, keyword := range s {
		fmt.Printf("%v : %v\n", i, keyword)
	}
}
