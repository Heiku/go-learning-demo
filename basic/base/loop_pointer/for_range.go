package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	// every loop, &value always change, point to the final value
	for index, value := range slice {
		//num := value
		myMap[index] = &value
	}
	fmt.Println("==== new map ====")
	prtMap(myMap)

	/**
	map[0]=3
	map[1]=3
	map[2]=3
	map[3]=3
	*/
}

func prtMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}
