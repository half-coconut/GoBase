package main

import "fmt"

// NewUser 如何初始化
func NewUser() {
	// 初始化结构体
	u := User{}
	fmt.Printf("%v \n", u)
	fmt.Printf("%#v \n", u)
}

type User struct {
	Name string
	Age  int
}
