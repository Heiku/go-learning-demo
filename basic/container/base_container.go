package main

// container means a set which can store element(key - value), such as map, array, slice, (string, chan)
// return value contain [0, len()]

const Size = 32

type Person struct {
	name string
	age  int
}

// array type
var a0 = [6]string{}
var a1 = [Size]int{}
var a3 = [16][]byte{}
var a4 = [100]Person{}

// slice type
var b0 = []bool{}
var b1 = []int64{}
var b2 = []map[int]bool{}
var b3 = []*int{}

// map type
var c0 = map[string]int{}
var c1 = map[int]bool{}
var c2 = map[int16][6]string{}       // value is [6]string
var c3 = map[bool][]string{}         // value is []string
var c4 = map[struct{ x int }]*int8{} // key is struct{x int}, value is *int8

// define and init
var d0 = [4]bool{false, true, true, false}
var d1 = []string{"break", "continue", "fallthrough"}
var d2 = map[string]int{"C": 1972, "Python": 1991, "Go": 2009}

// use ... make system infer the len of container
var base = [...]int{1, 3, 5, 7, 9}

func main() {

}
