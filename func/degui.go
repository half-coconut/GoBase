package main

import "fmt"

func main() {
	//fmt.Println(sum(100)
	// defer 函数延迟执行，函数执行到最后，会按照逆序执行
	// defer 应用场景 关闭操作
	//fmt.Println(1)
	//fmt.Println(2)
	//defer fmt.Println(3)
	//fmt.Println(4)
	//defer fmt.Println(5)
	//fmt.Println(6)
	//defer fmt.Println(7)
	//fmt.Println(8)

	// 函数本身就是一个数据类型
	// 函数不加括号，就是一个变量
	// sum() 加了括号，就是一个函数的调用
	//fmt.Printf("%T\n", sum)
	//fmt.Printf("%T\n", 101)
	//fmt.Printf("%T\n", "sum")

	//var f5 func(int) int
	//f5 = sum
	//fmt.Println(f5)
	//fmt.Println(sum)
	//fmt.Println(f5(100))

	// 匿名函数，本身可以调用自己
	func(a, b int) int {
		fmt.Println("这里是匿名函数", a+b)
		return a + b
	}(1, 2)

	r := func(a, b int) int {
		fmt.Println("这里是匿名函数", a+b)
		return a + b
	}(1, 2)
	fmt.Println(r)
}

func sum(n int) int {
	if n == 1 {
		return 1
	}

	return sum(n-1) + n
}
