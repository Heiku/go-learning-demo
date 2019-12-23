package main

// if struct with same type field, they can transform each other
type S0 struct {
	y int "foo"
	x bool
}

type S1 = struct { // S1 is a anonymous type
	x int "foo"
	y bool
}

type S2 = struct {
	x int "bar"
	y bool
}

type S3 S2
type S4 S3

var v0, v1, v2, v3, v4 = S0{}, S1{}, S2{}, S3{}, S4{}

func main() {
	v1 = S1(v2)
	v2 = S2(v1)

	v1 = S1(v3)
	v3 = S3(v1)

	v1 = S1(v4)
	v4 = S4(v1)

	v2 = v3
	v3 = v2
	v2 = v4
	v4 = v2

	v3 = S3(v4)
	v4 = S4(v3)
}
