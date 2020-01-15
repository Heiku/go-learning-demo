package main

import "fmt"

type Age int

// object compare
func (age Age) LargerThan(a Age) bool {
	return age > a
}

// pointer
func (age *Age) Increase() {
	*age++ // *pointer -> value ++
}

type FilterFunc func(in int) bool

func (ff FilterFunc) Filter(in int) bool {
	return ff(in)
}

type StringSet map[string]struct{}

func (ss StringSet) Add(key string) {
	ss[key] = struct{}{}
}
func (ss StringSet) remove(key string) {
	delete(ss, key)
}

type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

// will compile to
/*func Book.Pages(b Book) int{
	return b.pages
}*/

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

/*func (*Book).SetPages(b *Book, pages int){
	b.pages = pages
}*/

func main() {
	var book Book

	fmt.Printf("%T \n", book.pages)       // int
	fmt.Printf("%T \n", (&book).SetPages) // func(int)
	fmt.Printf("%T \n", (&book).Pages)    // func() int

	(&book).SetPages(123)
	book.SetPages(123) // syntactic sugar, it will transform the up method

	fmt.Println(book.Pages())
	fmt.Println((&book).Pages())
}
