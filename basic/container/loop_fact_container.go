package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
	}

	// array in loop is a duplicate
	persons := [2]Person{{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println(i, p)

		persons[1].name = "Jack"
		// didn't work, because index,value value just a duplicate of actual value
		p.age = 31
		// only can use person[i].age = 31 to change it
	}
	fmt.Println("person: ", persons)

	langs := map[struct{ dynamic, strong bool }]map[string]int{
		{true, false}: {
			"JavaScript": 1995,
		},
		{false, true}: {
			"Go": 2009,
		},
		{false, false}: {
			"C": 1972,
		},
	}
	m0 := map[*struct{ dynamic, strong bool }]*map[string]int{}
	for category, langInfo := range langs {
		m0[&category] = &langInfo
		// didn't work
		category.dynamic, category.strong = true, true
	}
	for category, langInfo := range langs {
		fmt.Println(category, langInfo)
	}

	m1 := map[struct{ dynamic, strong bool }]map[string]int{}
	// m0 just a duplicate
	for category, langInfo := range m0 {
		m1[*category] = *langInfo
	}
	for category, langInfo := range langs {
		fmt.Println(category, langInfo)
	}

	fmt.Println(len(m0), len(m1))
	fmt.Println(m1)

}
