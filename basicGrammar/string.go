package main

import (
	"fmt"
)

func String() {
	/**
	字符串类型
	*/
	var (
		apple  string = "苹果"
		peach         = "桃子"
		cherry        = "樱桃"
	)
	fmt.Printf("%T,%s,%s\n", apple, peach, cherry) // string,桃子,樱桃
	s1 := 'A'
	s2 := "A"
	fmt.Printf("%T,%d\n", s1, s1) // int32,65
	fmt.Printf("%T,%s\n", s2, s2) // string,A

	// 字符串连接
	fmt.Println(apple + s2 + peach + s2 + cherry) // 苹果A桃子A樱桃

	// 转义字符
	fmt.Println("hello\"pandan")
	fmt.Println("hello\npandan")
	fmt.Println("hello\tpandan")
}

func change() {
	/**
	转换数据类型
	*/
	a := 3
	b := 5.0
	c := float64(a)
	d := int64(b)
	//e := bool(a) // 整型不能转换成布尔类型
	fmt.Printf("%T,%d\n", a, a)
	fmt.Printf("%T,%f\n", b, b)
	fmt.Printf("%T,%f\n", c, c)
	fmt.Printf("%T,%d\n", d, d)
	//fmt.Printf("%T,%d\n", e, e)
}

func mathOperator() {
	/**
	算数运算符
	*/
	a := 3
	b := 10
	fmt.Println(a + b) // 13
	fmt.Println(a - b) // -7
	fmt.Println(a * b) // 30
	fmt.Println(a / b) // 0
	fmt.Println(a % b) // 3
	a++
	fmt.Println(a) // 4
	b--
	fmt.Println(b) // 9
}

func relationOperator() {
	/**
	关系运算符
	*/
	a := 3
	b := 10
	fmt.Println(a == b) // false
	fmt.Println(a >= b) // false
	fmt.Println(a <= b) // true
	fmt.Println(a != b) // true
	fmt.Println(a > b)  // false
	fmt.Println(a < b)  // true

	if a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a < b")
	}

}

func logicOperator() {
	/**
	逻辑运算符：&& || ！与 或 非
	*/
	a := true
	b := false
	if a && b {
		fmt.Println("a 和 b 都为 true")
	} else if a || b {
		fmt.Println("a 和 b 有一个为 true")
	} else {
		fmt.Println("a 和 b 都为 false")
	}
	fmt.Println(!a)
	fmt.Println(!b)
}

func bitwiseOperators() {
	/**
	位运算符：& | ^ &^ << 左移  >> 右移
	o false
	1 true
	*/
	var (
		a uint = 10
		b uint = 13
		c uint = 0
	)
	//c = a & b
	//fmt.Printf("%d，二进制：%b\n", a, a)
	//fmt.Printf("%d，二进制：%b\n", b, b)
	//fmt.Printf("%d，二进制：%b\n", c, c)
	/**
	a & b 同时满足
	10，二进制：1010
	13，二进制：1101
	8，二进制：1000
	*/
	//c = a | b
	//fmt.Printf("%d，二进制：%b\n", a, a)
	//fmt.Printf("%d，二进制：%b\n", b, b)
	//fmt.Printf("%d，二进制：%b\n", c, c)
	/**
	a | b 只要有一个满足即可
	10，二进制：1010
	13，二进制：1101
	15，二进制：1111
	*/
	//c = a ^ b
	//fmt.Printf("%d，二进制：%b\n", a, a)
	//fmt.Printf("%d，二进制：%b\n", b, b)
	//fmt.Printf("%d，二进制：%b\n", c, c)
	/**
	a | b 不同为1，相同为0
	10，二进制：1010
	13，二进制：1101
	7，二进制：111
	*/
	//c = a << 2
	//fmt.Printf("%d，二进制：%b\n", a, a)
	//fmt.Printf("%d，二进制：%b\n", b, b)
	//fmt.Printf("%d，二进制：%b\n", c, c)
	/**
	a 左移了2位
	10，二进制：1010
	13，二进制：1101
	40，二进制：101000
	*/
	c = b >> 2
	fmt.Printf("%d，二进制：%b\n", a, a)
	fmt.Printf("%d，二进制：%b\n", b, b)
	fmt.Printf("%d，二进制：%b\n", c, c)
	/**
	b 右移了2位
	10，二进制：1010
	13，二进制：1101
	3，二进制：11
	*/
}

func assignmentOperator() {
	/**
	赋值运算符(略)：
	=、+=、-=、*=、、=、%=、<<=、>>=、&=、^=、|=
	& 返回变量存储地址
	* 指针变量
	*/
}

func main() {
	/**
	input
	*/
	var (
		a int
		b float64
	)
	//fmt.Println() // 打印并换行
	//fmt.Printf()  // 格式化输出
	//fmt.Println() // 打印输出

	fmt.Println("请输入两个变量(1 整数 2 浮点数)：")
	fmt.Scan(&a, &b) // 接收输入
	fmt.Printf("请输入的两个变量是：%d, %f", a, b)
	//fmt.Scanf()  // 接收格式化输入
	//fmt.Scan()   // 接收输入
}
