package main

import "fmt"

func ifAndElse() {
	/**
	流程控制：顺序结构、选择结构、循环结构
	if switch for while
	*/
	var a = ""
	var b = ""
	password := "123456"
	fmt.Print("Please input your password: ")
	fmt.Scanln(&a)
	if a == password {
		fmt.Print("Please input your password again: ")
		fmt.Scanln(&b)
		if b == password {
			fmt.Print("Login Success! ")
		} else {
			fmt.Print("Sorry, Login Failure! The second password is NOT correct.")
		}
	} else {
		fmt.Print("Sorry, Login Failure! ")
	}
}

func switchFallthrough() {
	/**
	switch
	fallthrough: 贯穿，直通，很少用
	*/
	//for i := 0; i < 5; i++ {
	//	var score int64 = 0
	//	fmt.Print("Please input your score: ")
	//	fmt.Scanln(&score)
	//	fmt.Print("result is: ")
	//	switch score {
	//	case 90:
	//		fmt.Println("A")
	//	case 80:
	//		fmt.Println("B")
	//	case 70, 75, 60:
	//		fmt.Println("C")
	//	default:
	//		fmt.Println("You did NOT pass the exam! ")
	//	}
	//}

	a := true
	switch a {
	case true:
		fmt.Println("true")
		fallthrough
	case false:
		fmt.Println("false")
	}
	/**
	true
	false
	*/

}

func forLoop() {
	/**
	for 循环条件 {}:
		i 在 for后面
		i 在 for 前中后
	for {}: 不需要 i，相当于 while true 无限循环
	*/
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(sum)

	//for {
	//	fmt.Println("===")
	//}
}

func fivePrint() {
	/**
	打印如下
	* * * * *
	* * * * *
	* * * * *
	* * * * *
	* * * * *
	*/
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func ninetyNineMultiplicationTable() {
	/**
	打印99乘法表
	1*1=1
	2*1=2   2*2=4
	3*1=3   3*2=6   3*3=9
	4*1=4   4*2=8   4*3=12  4*4=16
	5*1=5   5*2=10  5*3=15  5*4=20  5*5=25
	6*1=6   6*2=12  6*3=18  6*4=24  6*5=30  6*6=36
	7*1=7   7*2=14  7*3=21  7*4=28  7*5=35  7*6=42  7*7=49
	8*1=8   8*2=16  8*3=24  8*4=32  8*5=40  8*6=48  8*7=56  8*8=64
	9*1=9   9*2=18  9*3=27  9*4=36  9*5=45  9*6=54  9*7=63  9*8=72  9*9=81
	*/
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	/**
	break
	continue
	*/
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("------------------")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
