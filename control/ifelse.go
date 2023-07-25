package main

func IfOnly(age int) {
	if age > 18 {
		println("成年了")
	}
}
func IfElse(age int) {
	if age >= 18 {
		println("成年了")
	} else {
		println("小孩子")
	}
}
func IfElseIf(age int) {
	if age >= 18 {
		println("成年了")
	} else if age >= 12 {
		println("青少年")
	} else {
		println("小孩子")
	}
}

func IfNewVariable(start int, end int) string {
	if distance := end - start; distance > 100 {
		return "太远了"
	} else if distance > 60 {
		return "有点远"
	} else {
		return "还可以"
	}
	// distance 仅在 if 语句里使用，方法内甚至都不能调用
	//println(distance)

}

func main() {
	//IfOnly(80)
	//IfElse(18)
	//IfElseIf(13)
	//println(IfNewVariable(20, 500))
	LoopV3()
	LoopV3()
}
