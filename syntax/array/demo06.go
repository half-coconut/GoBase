package main

import "fmt"

func main() {
	// 多维数组
	// [5]int
	// [5][4]int
	// [5][4][3]int

	// 3行4列
	arr := [3][4]int{
		{1, 2, 3, 4},    // 0
		{5, 6, 7, 8},    // 1
		{9, 10, 11, 12}, // 2
	}
	fmt.Println(arr[0][0])
	fmt.Println(arr[1][2])

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(arr[i][j])
		}
	}

	for i, v := range arr {
		fmt.Println(i, v)
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

}
