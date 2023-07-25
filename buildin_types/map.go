package main

// Map map 的遍历都是随机的，遍历两遍，输出的结果都不一样
func Map() {
	m1 := map[string]int{
		"key1": 123,
	}
	m1["hello"] = 999
	for k, _ := range m1 {
		println(k, m1[k])
	}

	// 容量，注意预付容量
	m2 := make(map[string]int, 12)
	m2["key2"] = 12
	for k, _ := range m2 {
		println(k, m2[k])
	}

	val, ok := m1["hello"]
	if ok {
		// 有这个键值对
		println("hello对应的值:", val)
	}
	val = m1["coconut"]
	println("coconut对应的值:", val)

	delete(m1, "hello")
	val = m1["hello"]
	println("hello对应的值:", val)
}
