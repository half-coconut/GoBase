package main

import "fmt"

func defineVar() {
	/**
	1 定义变量的几种方式，简写语法糖，打印内存地址
	*/

	var name0 string = "sss"
	var age0 int = 18

	var (
		name1 string
		age1  int
		addr1 string
	)

	// := 自动推导，申明初始化
	name := "coconut"
	age := 18
	addr := "中国"

	fmt.Println(name0, age0)
	fmt.Println(name, age, addr)
	fmt.Printf("%T,%T", name1, age1, addr1)
	fmt.Printf("%T,%T", name, age, addr)
	fmt.Println()
	fmt.Printf("num:%d，内存地址：%p", age, &age) // 取地址符 &
	age = 100
	fmt.Println()
	fmt.Printf("num:%d，内存地址：%p", age, &age) // 取地址符 &
	fmt.Println()
}

/*
*
2 全局变量和局部变量、变量的交换、匿名函数
*/
var name string = "coconut"

func test1() {
	/**
	变量交换
	*/
	a := 100
	b := 200
	a, b = b, a
	name := "zhangsan"
	fmt.Println(a, b, name)
}

func test2() (int, int) {
	fmt.Println(name)
	return 2, 3
}

func test3() {
	// 匿名函数 _
	a, _ := test2()
	test1()
	test2()
	fmt.Println(a, name)
}
