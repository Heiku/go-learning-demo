package main

import "fmt"

func test(a *[]int) {
	*a = append(*a, 4)
}

func main() {
	s := make([]int, 3, 4)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println(s)
	test(&s)
	// s = append(s, 4)
	fmt.Println(s)
}
