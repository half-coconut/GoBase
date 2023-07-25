package main

import "fmt"

// Array 数组，len 是长度，cap 是容量
func Array() {
	a1 := [3]int{9, 8, 7}
	fmt.Printf("a1: %v, len=%d, cap=%d \n", a1, len(a1), cap(a1))

	a2 := [3]int{9, 8}
	fmt.Printf("a2: %v, len=%d, cap=%d \n", a2, len(a2), cap(a2))

	a3 := [3]int{}
	fmt.Printf("a3: %v, len=%d, cap=%d \n", a3, len(a3), cap(a3))

	// a1: [9 8 7], len=3, cap=3
	// a2: [9 8 0], len=3, cap=3
	// a3: [0 0 0], len=3, cap=3
	println("\n", a2[1])
}
