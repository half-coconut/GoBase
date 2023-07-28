package main

import "fmt"

func add(a, b int) int {
	fmt.Printf("3/4 有两个参数且有一个返回值的函数:%d\n", a+b)
	return a + b
}

func printInfo() {
	fmt.Println("1 无参无返回值的函数")
}

func printStr(s string) {
	fmt.Println("2 有一个参数的函数:" + s)
}
func printNum(i int) {
	fmt.Printf("2 有一个参数的函数:%d\n", i)
}

func swap(a, b string) (string, string) {
	fmt.Printf("5 有多个返回值的函数:%s,%s\n", a, b)
	return b, a
}

func main() {
	/**
	函数的声明和调用：
	1 无参无返回值的函数
	2 有一个参数的函数
	3 有两个参数的函数
	4 有一个返回值的函数
	5 有多个返回值的函数
	===================
	1 无参无返回值的函数
	2 有一个参数的函数:coconut
	3/4 有两个参数且有一个返回值的函数:3
	2 有一个参数的函数:3
	5 有多个返回值的函数:xiaoer,Pandan
	Pandan xiaoer
	*/

	printInfo()
	printStr("coconut")
	printNum(add(1, 2))
	a, _ := swap("xiaoer", "Pandan")
	b, c := swap("xiaoer", "Pandan")
	fmt.Println(a)
	fmt.Println(b, c)
}
