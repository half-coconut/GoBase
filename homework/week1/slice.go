package main

import (
	"errors"
	"fmt"
)

var ErrIndexOutOfRange = errors.New("下标超出范围")

// DeleteAt 删除指定位置的元素
// 如果下标不是合法的下标，返回 ErrIndexOutOfRange
// 要求一：能够实现删除操作
func DeleteAt(s []int, idx int) ([]int, error) {
	if idx < 0 || idx >= len(s) {
		return nil, ErrIndexOutOfRange
	}
	res := make([]int, 0, len(s)+1)
	for i := 0; i < idx; i++ {
		res = append(res, s[i])
	}
	for i := idx + 1; i < len(s); i++ {
		res = append(res, s[i])
	}
	return res, nil
}

// DeleteAtV3
// 要求三：改造为泛型方法
func DeleteAtV3[T any](s []T, idx int) ([]T, error) {
	if idx < 0 || idx >= len(s) {
		return nil, ErrIndexOutOfRange
	}
	res := make([]T, 0, len(s)+1)
	for i := 0; i < idx; i++ {
		res = append(res, s[i])
	}
	for i := idx + 1; i < len(s); i++ {
		res = append(res, s[i])
	}
	return res, nil
}

func main() {
	slice := []int{9, 8, 7, 6, 5}
	res, e := DeleteAt(slice, 4)
	fmt.Printf("res: %+v \n", res)
	fmt.Printf("e: %+v \n", e)

	s2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	res2, e2 := DeleteAtV3(s2, 4)
	fmt.Printf("res2: %+v \n", res2)
	fmt.Printf("e2: %+v \n", e2)
}
