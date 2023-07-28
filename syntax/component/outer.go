package main

// Inner 组合，很像其他语言的继承
// 但是组合不是继承，没有多态
type Inner struct {
}

func (i Inner) DoSomething() {
	println("这是 Inner")
}
func (o Outer) DoSomethingV2() {
	println("这是 Outer")
}
func (o OuterV1) DoSomething() {
	println("这是 OuterV1")
}

func (i Inner) SayHello() {
	println("hello", i.Name())
}

func (i Inner) Name() string {
	return "inner"
}

type Outer struct {
	Inner // 大部分使用这种
}

func (o Outer) Name() string {
	return "outer"
}

type OuterV1 struct {
	Inner // 大部分使用这种
}
type OuterPtr struct {
	*Inner // 了解，如果第三方使用指针，就用指针
}

type OOOuter struct {
	Outer
}

func UseInner() {
	var o Outer
	// 组合之后，就可以调用方法了
	o.DoSomething()
	o.Inner.DoSomething()
	o.DoSomethingV2()

	var op *OuterPtr
	op.DoSomething()
	op.Inner.DoSomething()
	op.DoSomething()

	var oo OOOuter
	oo.DoSomethingV2()
	oo.DoSomething()
	oo.Inner.DoSomething()
	oo.Outer.Inner.DoSomething()
	oo.Outer.DoSomethingV2()

	o1 := Outer{
		Inner: Inner{},
	}
	op1 := OuterPtr{
		Inner: &Inner{},
	}
	o1.Inner.DoSomething()
	o1.DoSomethingV2()
	o1.DoSomething()
	op1.Inner.DoSomething()
	op1.DoSomething()
}

func main() {
	//var o1 OuterV1
	//o1.Inner.DoSomething()
	//o1.DoSomething()

	var o Outer
	o.SayHello() // hello inner

}
