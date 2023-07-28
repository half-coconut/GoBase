package main

import "fmt"

func main() {
	// 数组是值传递，是拷贝，不影响原数据
	num := 10
	fmt.Printf("%T\n", num)

	arr1 := [3]int{1, 2, 3}
	arr2 := [2]string{"apple", "cat"}
	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)
}
