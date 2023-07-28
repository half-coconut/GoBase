package genetics

// List 泛型，T 类型参数，名字叫 T，约束是 any，等于没有约束
type List[T any] interface {
	Add(idx int, t T)
	Append(t T)
}

func UseList() {
	var l List[int]
	l.Append(10)
	l.Append(12.3) // 但是，这里其实是用不了
}
