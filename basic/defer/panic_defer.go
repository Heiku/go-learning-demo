package main

import "fmt"

func main() {
	// defer func value valuated before push defer stack
	// so f = func(){} can't stop panic causing

	defer fmt.Println("can invoked")
	var f func() //	 f == nil
	defer f()    // will cause panic

	fmt.Println("can invoked too")
	f = func() {} // can't stop causing panic
}
