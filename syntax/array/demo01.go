package main

import "fmt"

func main() {
	// 数组是相同类型数据的有序集合
	// 数组一旦定义后，大小不能改变
	var nums [5]int

	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4
	nums[4] = 5
	fmt.Printf("%T\n", nums)
	fmt.Println(nums)
	fmt.Println(nums[3])
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	nums[0] = 100
	fmt.Println(nums)

}
