package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Sum[T Number](vals ...T) T {
	var res T
	for _, val := range vals {
		res = res + val
	}
	return res
}

type Integer int

type Number interface {
	// ~int 指 int 以及 int 的衍生类型
	~int | int64 | float64 | float32 | int32 | byte | uint
}

type MyMarshal struct {
}

func (m *MyMarshal) MarshalJSON() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func ReleaseResource[R json.Marshaler](r R) {
	r.MarshalJSON()
}

func Max[T Number](vals ...T) (T, error) {
	if len(vals) == 0 {
		var t T
		return t, errors.New("下标不对")
	}

	res := vals[0]
	for i := 1; i < len(vals); i++ {
		if res < vals[i] {
			res = vals[i]
		}
	}
	return res, nil
}

func Min[T Number](vals ...T) (T, error) {
	if len(vals) == 0 {
		var t T
		return t, errors.New("下标不对")
	}

	res := vals[0]
	for i := 1; i < len(vals); i++ {
		if res > vals[i] {
			res = vals[i]
		}
	}
	return res, nil
}

// AddSlice 实现在特定的位置，插入某个值
func AddSlice[T any](slice []T, idx int, val T) ([]T, error) {
	// 如果 idx 是负数或超过了 slice 的界限
	if idx < 0 || idx >= len(slice) {
		return nil, errors.New("下标出错")
	}

	res := make([]T, 0, len(slice)+1)
	for i := 0; i < idx; i++ {
		res = append(res, slice[i])
	}
	slice[idx] = val
	for i := idx; i < len(slice); i++ {
		res = append(res, slice[i])
	}
	return res, nil
}

func main() {
	//println(Sum(1, 2, 3))
	//println(Sum[int](1, 2, 3))
	//println(Sum[Integer](1, 2, 3))
	//println(Sum[float64](1.1, 2.1, 3.1))
	//println(Sum[float32](1.1, 2.1, 3.1))
	//var j MyMarshal
	//ReleaseResource[*MyMarshal](&j)
	slice := []int{9, 8, 7}
	fmt.Printf("slice: %v, len=%d, cap=%d \n", slice, len(slice), cap(slice))
	res, e := AddSlice[int](slice, 1, 10)
	fmt.Printf("res: %+v \n", res)
	fmt.Printf("e: %+v \n", e)
}
