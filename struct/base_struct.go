package main

import "fmt"

type Book struct {
	title  string `json:"title"`
	author string `json:"author,omitempty"`
	pages  int    `json:"page,omitempty"`
}

func main() {
	// direct newInstance
	book := Book{"golang 101", "laomo", 256}
	fmt.Println(book)

	// newInstance with field
	book = Book{
		title:  "golang 101",
		author: "laomo",
		pages:  300,
	}
	book = Book{}
	book = Book{author: "laomo"}

	// use chooser(getter/setter) change field value
	book2 := Book{}
	book2.author = "gopher"
	book2.pages = 301
	fmt.Println(book2)

}
