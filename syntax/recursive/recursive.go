package main

func Recursive(n int) {
	// 递归注意要有退出条件，否则会有 stack overflow
	if n > 10 {
		return
	}
	Recursive(n + 1)
}

func A() {
	B()
}
func B() {
	C()
}
func C() {
	A()
}
