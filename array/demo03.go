package main

import "fmt"

func main() {
	// fori
	for i := 0; i < 4; i++ {

	}
	// arr1.for
	arr1 := [4]int{5, 6, 7, 8}
	for i, i2 := range arr1 {
		fmt.Println("index: ", i)
		fmt.Println("value: ", i2)
	}

}
