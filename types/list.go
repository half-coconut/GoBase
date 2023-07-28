package main

// List 接口定义
type List interface {
	Add(idx int, val any) error
	Append(val any)
	Delete(val any) (any, error)
	toSlice() ([]any, error) // 私有的方法，不常见
}
