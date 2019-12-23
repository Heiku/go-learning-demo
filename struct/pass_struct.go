package main

import "fmt"

type Book struct {
	title  string
	author string
	pages  int64
}

func f(book Book) {
	book.pages = 400
	book.title = "golang up"
}

func F(book *Book) {
	book.pages = 400
	book.title = "golang up"
}

func main() {
	book := Book{
		title:  "golang 101",
		author: "laomo",
		pages:  256,
	}

	// is same as value, change just happen on param duplicate
	fmt.Println(book)
	f(book)
	fmt.Println(book)

	// use pointer to change it
	F(&book)
	fmt.Println(book)

	b := &book
	b.pages = 500 // pointer also can change field
	fmt.Println(book)
}
