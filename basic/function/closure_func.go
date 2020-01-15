package main

import "fmt"

func main() {
	// closure
	isMultipleOfX := func(x int) func(int) bool {
		return func(n int) bool {
			return n%x == 0
		}
	}
	var isMultipleOf3 = isMultipleOfX(3)
	var isMultipleOf5 = isMultipleOfX(5)

	fmt.Println(isMultipleOf3(6))
	fmt.Println(isMultipleOf3(8))
	fmt.Println(isMultipleOf5(10))
	fmt.Println(isMultipleOf5(12))

	isMultipleOf15 := func(n int) bool {
		return isMultipleOf3(n) && isMultipleOf5(n)
	}
	fmt.Println(isMultipleOf15(32))
	fmt.Println(isMultipleOf15(60))
}
