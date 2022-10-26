package main

import "fmt"

func forString() {
	str := "hello, coconut!"
	fmt.Println(str)
	// 获取字符串长度
	fmt.Println("string's length: ", len(str))
	// 获取指定的字节,%c char
	fmt.Printf("the first byte: %c\n", str[0])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
}

func main() {
	/**
	for range 遍历数组、切片等
	*/
	str := "hello coconut!"
	for i, v := range str {
		fmt.Printf("%d,%c\n", i, v)
	}
}
