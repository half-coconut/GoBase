package main

import "fmt"

func max(a, b int) int {
	/**
	形参和实参
	*/
	// 这里 a 和 b 为形参
	if a > b {
		return a
	} else {
		return b
	}
}

func getSum(nums ...int) int {
	/**
	可变参数 ...xxx
	位置一定要放在最后
	*/
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func params(arr [4]int) {
	/**
	值类型的数据：操作的是数据本身 int、string、bool、float64、array...
	引用类型的数据：操作的是数据的地址 slice、map、chan...
	值传递：传递的是数据的副本，修改数据，对于原始的数据没有影响
	值类型的数据，默认都是值传递，基础类型，array、struct
	定义一个数组 [个数]类型
	*/

	fmt.Println(arr)
}
func update(arr [4]int) {
	fmt.Println(arr)
	arr[0] = 100
	fmt.Println(arr)
}

func update2(s2 []int) {
	fmt.Println(s2)
	s2[0] = 100
	fmt.Println(s2)
}

func main() {
	// 这里就是 实参
	//fmt.Println(max(45, 15))
	//fmt.Println("sum: ", getSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	//arr := [4]int{1, 2, 3, 4}
	//update(arr)
	//fmt.Println(arr)

	// 切片，可以扩容的数组
	s1 := []int{1, 2, 3, 4}
	update2(s1)
	fmt.Println(s1)
}
