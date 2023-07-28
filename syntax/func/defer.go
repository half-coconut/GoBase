package main

// Defer 的特点：先进后出，后进先出(同 stack)
func Defer() {
	defer func() {
		println("第一个 defer")
	}()

	defer func() {
		println("第二个 defer")
	}()
}

// DeferClosure defer 闭包
func DeferClosure() {
	i := 0
	defer func() {
		println(i)
	}()
	i = 1
}

// DeferClosureV2 这里传参i 作为参数，初始化为 0 时，已经传入
func DeferClosureV2() {
	i := 0
	defer func(i int) {
		println(i)
	}(i)
	i = 1
}

// DeferClosureV3 这里是 go 函数有返回值名称的特性，这里的 a 因为是返回值，所以赋值成功，defer 返回的是 a=1
func DeferClosureV3() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

// defer 练习题

// DeferClosureLoopV1 全都是 10
func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

// DeferClosureLoopV2 按照"先进后出"原则，倒序输出 for 循环的 i 值
func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println(val)
		}(i)
	}
}

// DeferClosureLoopV3 同 DeferClosureLoopV2 的输出，相当于 j 为变量，i传值给 j
func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}

func demoDefer() {
	//Defer()
	//DeferClosure()
	//DeferClosureV2()
	//println(DeferClosureV3())
	//DeferClosureLoopV1()
	//DeferClosureLoopV2()
	DeferClosureLoopV3()
}
