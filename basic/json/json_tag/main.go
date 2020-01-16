package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	/*Name string `name`
	Age int	`age`*/

	Name string `json:"name" bson:"b_name"`
	Age  int    `json:"age" bson:"b_age"`
}

// parse by struct "tag"
func main() {

	// deserialize
	var u User
	h := `{"name":"张三", "age": 15}`
	err := json.Unmarshal([]byte(h), &u)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u)
	}

	// serialize
	newJson, err := json.Marshal(u)
	fmt.Println((string(newJson)))

	// reflect
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Tag)
		fmt.Println(t.Field(i).Tag.Get("json"))
		fmt.Println(t.Field(i).Tag.Get("bson"))
	}
}
