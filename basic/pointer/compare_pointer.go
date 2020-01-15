package main

// only same type pointer can be compared

func main() {
	type MyInt int64
	type Ta *int64
	type Tb *MyInt

	// 4 difference type pointer
	var pa0 Ta
	var pa1 *int64
	var pb0 Tb
	var pb1 *MyInt

	// compile successfully
	_ = pa0 == pa1
	_ = pb0 == pb1
	_ = pa0 == nil
	_ = pa1 == nil
	_ = pb0 == nil
	_ = pa1 == nil

	// compile wrong
	/**
	_ = pa0 == pb0
	_ = pa1 == pb1
	_ = pa0 == Tb(nil)
	*/
}
