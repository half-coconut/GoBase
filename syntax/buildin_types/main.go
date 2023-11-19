package main

import "fmt"

func main() {
	//Array()
	//Slice()
	//SubSlice()
	//ShareSlice()
	//Map()

	a1 := []int{5, 7, 7, 8, 8, 10}
	fmt.Printf("a1: %v, len=%d, cap=%d \n", a1, len(a1), cap(a1))

	a2 := make([]int, 0)
	a2 = append(a2, 3)
	a2 = append(a2, 4)
	fmt.Printf("a2: %v, len=%d, cap=%d \n", a2, len(a2), cap(a2))
	fmt.Printf("a2: %v", searchRange(a1, 8))
}
