package main

// LinkedList 结构体
type LinkedList struct {
	head *node // 指针
	tail *node

	// 这就是包外可访问
	Len int
}

type node struct {
}
