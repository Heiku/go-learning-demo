package main

import (
	"fmt"
)

func main() {
	values := []interface{}{
		456, "abc", true, 0.33, int32(789), []int{1, 2, 3}, map[int]bool{}, nil,
	}

	for _, x := range values {
		switch x.(type) {
		case []int:
			fmt.Println("int slice: ", x)
		case string:
			fmt.Println("string: ", x)
		case int, float64, int32:
			fmt.Println("number: ", x)
		case nil:
			fmt.Println("nil", x)
		default:
			fmt.Println("others: ", x)
		}
	}
}
