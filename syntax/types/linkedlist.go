package main

import "time"

// LinkedList 结构体，使用结构体实现接口
type LinkedList struct {
	head *node // 指针
	tail *node

	// 这就是包外可访问
	Len int

	CreateTime time.Time
}

func (l *LinkedList) Add(idx int, val any) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Append(val any) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Delete(val any) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) toSlice() ([]any, error) {
	//TODO implement me
	panic("implement me")
}

//func (l LinkedList) Add(idx int, val any) {
//
//}
//
//// AddV1 (l *LinkedList) 方法接收器 receiver
//func (l *LinkedList) AddV1(idx int, val any) {
//
//}

type node struct {
	prev *node // 自引用，改成指针(指针是固定长度的)，好计算使用空间
	next *node
	//next node  自引用，不用指针会编译错误
}

// 指针，本质上就是一个内存地址，* 表示指针，& 取地址的符号
// 如果声明了一个指针，但是没有赋值，那么它是 nil
