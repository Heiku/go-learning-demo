package main

import "fmt"

func main() {
	type Ta []int
	type Tb []int
	dest := Ta{1, 2, 3}
	src := Tb{5, 6, 7, 8, 9}

	n := copy(dest, src)   // copy() return moving amount of element, moving count depend on the len of dest array
	fmt.Println(dest, src) // copy Tb{5, 6, 7} into desc and cover {1, 2, 3}

	n = copy(dest[1:], dest)
	fmt.Println(n, dest)

	a := [4]int{}
	n = copy(a[:], src)
	fmt.Println(n, a)
	n = copy(a[:], a[2:]) // a[2:] = {7, 8}
	fmt.Println(n, a)
}
