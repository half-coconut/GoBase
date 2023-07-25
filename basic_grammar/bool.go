package main

import "fmt"

func main() {
	/**
	bool 不赋值时，默认为 false
	*/
	var isFlag bool
	var isTrue = true
	fmt.Println(isFlag, isTrue)

	// %T类型 %t 值
	fmt.Printf("%T,%d\n", isFlag, isTrue)
}
