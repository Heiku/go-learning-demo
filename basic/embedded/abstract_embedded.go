package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) PrintName() {
	fmt.Println("Name: ", p.Name)
}

func (p *Person) SetAge(age int) {
	p.Age = age
}
func (p Person) SetAge2(age int) {
	p.Age = age
}

type Singer struct {
	Person // expend Person method
	works  []string
}

func main() {
	var gaga = Singer{Person: Person{"GaGa", 30}}

	// type Singer can invoke Person method
	gaga.PrintName()
	gaga.Name = "Lady GaGa" // also can use Person's fields		gaga.Person.Name(Person can be omitted)
	gaga.SetAge2(31)        // didn't work, need pass the pointer
	gaga.PrintName()

	fmt.Println(gaga.Age)

	// reflect sub type Singer
	t := reflect.TypeOf(Singer{})
	fmt.Println(t, " has ", t.NumField(), " fields: ")
	for i := 0; i < t.NumField(); i++ {
		fmt.Print(" field#", i, " : ", t.Field(i).Name, " \n")
	}

	fmt.Println(t, " has ", t.NumMethod(), " methods: ")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Print(" method#", i, " : ", t.Method(i).Name, " \n")
	}

	pt := reflect.TypeOf(&Singer{})
	fmt.Println(pt, " has ", pt.NumMethod(), " methods: ")
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Print(" method#", i, " : ", t.Method(i).Name, " \n")
	}
}
