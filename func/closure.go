package main

import "fmt"

// Closure 闭包函数
func Closure(name string) func() string {
	return func() string {
		return "hello, " + name
	}
}

func Closure2() func() string {
	name := "coconut"
	age := 18
	return func() string {
		return fmt.Sprintf("hello, %s,%d", name, age)
	}
}

func Closure3() func() int {
	age := 0
	fmt.Printf("out: %p", &age)
	return func() int {
		fmt.Printf(" before: %p", &age)
		age++
		fmt.Printf(", after: %p", &age)
		return age
	}
}
