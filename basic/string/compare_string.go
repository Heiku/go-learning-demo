package main

import (
	"fmt"
	"time"
)

func main() {
	bs := make([]byte, 1<<26)

	// s0 -> bs's deep copy byte[]
	// s1 -> bs's deep copy byte[]
	s0 := string(bs)
	s1 := string(bs)

	// s2 & s1 use same copy byte[]
	s2 := s1
	s2 = "nono"

	// if comparing with same pointer, so just compare the len of _string, O(1)
	// if comparing with difference pointer, loop compare the byte in []byte, O(n)
	//duration for (s0 == s1):  6.9847ms
	//duration for (s1 == s2):  0s
	startTime := time.Now()
	_ = s0 == s1
	duration := time.Now().Sub(startTime)
	fmt.Println("duration for (s0 == s1): ", duration)

	startTime = time.Now()
	_ = s1 == s2
	duration = time.Now().Sub(startTime)
	fmt.Println("duration for (s1 == s2): ", duration)
}
