package main

func LoopV1() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}
}

func LoopV2() {
	for true {
		println("死循环")
	}
}

// LoopV3 注意 map 的 key 值每次读取都是乱序的
func LoopV3() {
	m := map[string]int{
		"key1": 100,
		"key2": 102,
	}
	for k, _ := range m {
		println(k, m[k])
	}
}

// 注意：for 循环不要对迭代参数取地址，会出错。
