package main

func YourNameV1(name string, aliases ...string) {
	println("v1: ", name, aliases)
}
func YourNameV2(name string, aliases ...any) {
	println("v2: ", name, aliases)
}

func CallYourName() {
	YourNameV1("coconut")
	YourNameV1("coconut", "switch on")
	YourNameV1("coconut", "switch on", "Panda", "cat")
	aliases := []string{"coco", "胖蛋", "三胖"}
	YourNameV1("小二", aliases...)
	YourNameV2("小二", aliases) // 注意这里是个坑
}

func main() {
	CallYourName()
}
