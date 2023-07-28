package main

import (
	"fmt"
	"time"
)

// Integer 是 int 的衍生类型
// 衍生类型是一个全新的类型，注意 Int 实现了某个接口，不等于 Integer 也实现了某个接口
type Integer int

func UseInt() {
	i1 := 10
	i2 := Integer(i1)
	var i3 Integer = 13
	println(i2, i3)
}

type Fish struct {
	Name   string
	length int
}

func (f Fish) Swim() {
	println("Fish is swimming")

}

type FakeFish Fish

func UseFish() {
	f1 := Fish{}
	f2 := FakeFish{}
	f1.Swim()
	//f2. 没有 Swim() 方法
	f2.Name = "Tom" // 改完之后，只改了自己的
	println("f1:", f1.Name)
	println("f2:", f2.Name)
	f2.length = 10 // 可以访问字段，但是访问不了方法，因为 FakeFish 是个全新的类型
	println("f1:", f1.length)
	println("f2:", f2.length)

	var y Yu
	y.length = 4
	y.Name = "金鱼"
	y.Swim()
	fmt.Printf("y: %+v \n", y)
}

// MyTime 衍生类型，使用第三方库时，改不了原本的方法，就用衍生类型
type MyTime time.Time

func (m MyTime) MyFunc() {

}

// Yu 是 Fish 的类型别名，使用场景：向后兼容
type Yu = Fish
