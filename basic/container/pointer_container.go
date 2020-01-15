package main

import "fmt"

// sometimes. we can use pointer instead of array variable, it can be more efficiently

func main() {
	pa := &[5]int{2, 3, 5, 7, 11}
	s := pa[1:3]
	fmt.Println(s)

	pa = nil
	s = pa[0:0]
}
