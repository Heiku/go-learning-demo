package main

import "fmt"

func main() {
	m0 := map[int]int{0: 7, 1: 8, 2: 9} // m0 -> map
	m1 := m0                            // m1 -> map, they share base container
	m1[0] = 2
	fmt.Println(m0, m1)

	s0 := []int{7, 8, 9}
	s1 := s0 // same as map, s0 s1 share same memory space
	s1[0] = 2
	fmt.Println(s0, s1)

	a0 := [...]int{7, 8, 9}
	a1 := a0
	a1[0] = 2
	fmt.Println(a0, a1)

	// delete
	m := map[string]int{"Go": 2007}
	m["C"] = 1972
	m["Java"] = 1995
	fmt.Println(m)

	m["Go"] = 2009
	delete(m, "Java")
	fmt.Println(m)

	// every call append(), it would allocate new memory space for new slice, and the capacity is double than before
	// so append() will allocate more memory(cap * 2) to reduce call allocate()
	fmt.Println(s0)
	s2 := append(s0, 10)
	fmt.Println(s2)
}
