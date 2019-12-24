package main

import "fmt"

type Aboutable interface {
	About() string
}

type Book struct {
	name string
}

func (b *Book) About() string {
	return "Book: " + b.name
}

func main() {
	// Aboutable(Book)
	var a Aboutable = &Book{"golang 101"}
	fmt.Println(a)

	// interface{} contain any interface method
	var i interface{} = &Book{"Rust 101"}
	fmt.Println(i)

	i = a
	fmt.Println(i)
}
