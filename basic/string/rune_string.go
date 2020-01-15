package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func Runes2Bytes(rs []rune) []byte {
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r)
	}

	n, bs := 0, make([]byte, n)
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
	}
	return bs
}

func main() {
	s := "hello 你好"

	// len would return the len of byte array
	// str is []byte, and a chinese character hold 3 byte, so is 5 + 1 + 2 * 3 = 12
	fmt.Println("len(str): ", len(s))

	// if we need the character length?
	// use rune, rune = int32, using for utf8 or unicode character
	fmt.Println("RuneCountInString: ", utf8.RuneCountInString(s))

	// return len of unicode
	fmt.Println("rune: ", len([]rune(s)))

	// transform string to rune or byte
	bs := []byte(s)
	s = string(bs)

	rs := []rune(s)
	s = string(rs)

	rs = bytes.Runes(bs)
	bs = Runes2Bytes(rs)

	s = "eि́Ꮻaπ囧"
	for i, rn := range s {

		// i show the byte interval
		fmt.Printf("%2v: 0x%x %v \n", i, rn, string(rn))
	}
	fmt.Println(len(s))
}
