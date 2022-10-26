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

func main() {
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
