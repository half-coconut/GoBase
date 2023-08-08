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
		return nil, fmt.Errorf("ekit: %w，下标超出范围，长度 %d，下标 %d", ErrIndexOutOfRange, len(s), idx)
	}
	for i := idx; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}

	return s[:len(s)-1], nil
}

// DeleteAtV3
// 要求三：改造为泛型方法
func DeleteAtV3[T any](s []T, idx int) ([]T, error) {
	if idx < 0 || idx >= len(s) {
		return nil, fmt.Errorf("ekit: %w，下标超出范围，长度 %d，下标 %d", ErrIndexOutOfRange, len(s), idx)
	}
	for i := idx; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}

	return s[:len(s)-1], nil
}

// Shrink 缩容
func Shrink[T any](s []T) []T {
	c, l := cap(s), len(s)
	n, changed := calCapacity(c, l)
	if !changed {
		return s
	}
	src := make([]T, 0, n)
	src = append(src, s...)
	return src
}

func calCapacity(c, l int) (int, bool) {
	// 容量小于 64，不考虑缩容
	if c <= 64 {
		return c, false
	}
	// 如果大于 2048，但是元素不足一半
	// 降低为 0.625，也就是 5/8
	// 比一半多一点，和正向扩容的 1.25 倍相呼应
	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}
	// 如果在 2048 以内，并且元素不足 1/4，那么直接缩减为一半
	if c <= 2048 && (c/l >= 4) {
		return c / 2, true
	}
	// 整个实现的核心是希望在后续少触发扩容的前提下，一次性释放尽可能多的内存
	return c, false
}

func main() {
	slice := []int{9, 8, 7, 6, 5}
	res, e := DeleteAt(slice, 1)
	fmt.Printf("res: %+v \n", res)
	fmt.Printf("e: %+v \n", e)

	s2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	res2, e2 := DeleteAtV3(s2, 2)
	fmt.Printf("res2: %+v \n", res2)
	fmt.Printf("e2: %+v \n", e2)
}
