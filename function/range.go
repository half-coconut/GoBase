package main

import "fmt"

var num int = 101

func main() {
	temp := 100
	num := 20
	if b := 1; b <= 10 {
		temp := 50
		fmt.Println(temp)
		fmt.Println(b)
		fmt.Println(num)
	}
	fmt.Println(temp)
	fmt.Println(num)
	f1()
}

func f1() {
	num := 10
	fmt.Println(num)
}
func f2() {
	fmt.Printf("")
}
