package main

func Functional() string {
	println("hello functional")
	return "hello"
}

func Functional2(age int) string {
	return "hello 2"
}

func useFunctional() {
	// 注意这里是 Functional，不带括号
	myFunc := Functional
	myFunc()
	// 没有方法重载，需要传参，使用新的方法(带参数)，新的变量接收
	myFunc2 := Functional2
	myFunc2(18)
}

func Functional3() {
	// 局部方法，新定义一个方法，赋值给了 fc
	// 作用域，只能在本方法内部使用，不常用
	fc := func() string {
		return "hello"
	}
	fc()
}

// 方法作为返回值：fuctional4 返回一个，返回 string 的无参数方法
func fuctional4() func() string {
	return func() string {
		return "hello, this is func"
	}
}

// 匿名方法立刻发起调用：直接调用，返回 string 的无参数方法
func fuctional5() {
	hello := func() string {
		return "hello, world"
	}()
	println(hello)
}

func demo3() {
	//println(Global)
	//println(internal)
	//a := fuctional4()
	//println(a())
	//fuctional5()
	//c := Closure2()
	//println(c())
	getAge := Closure3()
	println(", age:", getAge())
	println(", age:", getAge())
	println(", age:", getAge())
	println(", age:", getAge())
	println(", age:", getAge())
}
