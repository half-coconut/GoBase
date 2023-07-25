package main

import "fmt"

func main() {
	// 常规数组的初始化方式
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	// 快速定义数组
	arr2 := [5]int{6, 7, 8, 9, 10}
	fmt.Println(arr2)

	// ...不确定数组有多大
	arr3 := [...]string{"apple", "cat", "dog"}
	fmt.Println(arr3)

	arr4 := [5]int{1: 100, 4: 400} // int 默认值为0
	fmt.Println(arr4)

}
