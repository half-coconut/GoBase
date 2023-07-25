package main

import "fmt"

func demo1() {
	/**
	1 常量，不可修改
	*/
	const URL string = "www.baidu.com" // 显示定义
	const NUM = 12                     // 隐式定义
	fmt.Println(URL, NUM)

	const a, b, c = 3.14, "coconut", false // 同时定义多个值
	fmt.Println(a, b, c)
}

func main() {
	/**
	2 iota 特殊常量，常量计数器
	*/
	//const (
	//	a = iota
	//	b = iota
	//	c = iota
	//)
	const (
		a = iota
		b
		c
		d = "haha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	demo1()
	// 0 1 2 haha haha 100 100 7 8
}
