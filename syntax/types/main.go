package main

func main() {
	//NewUser() // illegal types for operand: print
	//ChangeUser()
	//UseInt()
	UseFish()

	//var l List
	//l = &LinkedList{}
	//l = &ArrayList{}
	//println(l)
}

func DoSomething(l List) {
	// l 为 List，可以是 LinkedList 也可以是 ArrayList，方法随便传，不在意底层实现
	l.Append(12.3)
	l.Append(10)
	l.Append("string")
}
