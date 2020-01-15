package main

type I interface {
	m()
}

type T struct {
	I
}

func main() {
	var t T
	var i = &t // i
	t.I = i

	t.m() // t.m() -> t.I.m()
	// -> I.m() & t.I = i
	// -> i.m() && i = &t
	// -> t.m() ....
}
