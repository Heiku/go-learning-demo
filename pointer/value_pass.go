package main

import "fmt"

// func calling is a process of copying param duplicate
func double(x int) {
	x += x // change is a duplicate of param x
}

func main() {
	var a = 3
	double(a)
	fmt.Printf("a = %d", a)
}
