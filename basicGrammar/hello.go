package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {
	// She said: "Hello Go!"
	println("She said: \"Hello Go!\"")
	println(`
可以换行
再来一行
`)
	println("hello" + "go")
	var a byte = 'a'
	println(fmt.Sprintf("%c", a))
	// 统计中文字数的个数
	println(utf8.RuneCountInString("你好"))
	println("int 最大值", math.MaxInt)
	println("int 最小值", math.MinInt)
	println("int8 最大值", math.MaxInt8)
	println("int8 最小值", math.MinInt8)
	println("int64 最大值", math.MaxInt64)
	println("int64 最小值", math.MinInt64)
	println("float64 最大值", math.MaxFloat64)
	println("float64 最小正数", math.SmallestNonzeroFloat64)
	println("float32 最大值", math.MaxFloat32)
	println("float32 最小正数", math.SmallestNonzeroFloat32)
}
