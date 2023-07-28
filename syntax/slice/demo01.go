package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	// 定义切片，长度是可变的
	// 切片是引用类型
	var s1 []int
	fmt.Println(s1)
	if s1 == nil {
		fmt.Println("切片为空")
	}

	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println(s2)
	fmt.Printf("%T,%T\n", s1, arr) // []int,[4]int

}
