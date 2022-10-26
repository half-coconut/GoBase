package main

import "fmt"

func main() {
	// byte 类似于 uint8
	var num byte = 23
	fmt.Printf("%T,%d\n", num, num) //uint8,23
	// 定义一个整形
	var age int = 18
	fmt.Printf("%T,%d\n", age, age) //int,18

	// 定义一个浮点型，float 打印是 %f，可以加 .xx 代表保留几位小数
	// go 里默认的 float 为 float64
	var money float64 = 3.14
	fmt.Printf("%T,%f\n", money, money) // float64,3.140000

	fmt.Printf("%T,%.2f", money, money) // float64,3.14

}
