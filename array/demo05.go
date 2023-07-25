package main

import "fmt"

func main() {
	// 数组的排序：冒泡排序、选择排序、希尔排序、堆排序、快速排序...

	arr := [5]int{5, 4, 2, 3, 1}
	sort(arr, "asc")
	sort(arr, "desc")
	//  [4 2 3 1 5]
	//	[2 3 1 4 5]
	//	[2 1 3 4 5]
	//	[1 2 3 4 5]
	//	[1 2 3 4 5]

}
func sort(arr [5]int, c string) [5]int {
	// [4 2 3 1 5]
	for j := 0; j < len(arr); j++ {
		for i := 0; i < len(arr)-j-1; i++ {
			// 这里控制升序，或者降序
			if c == "asc" {
				if arr[i] > arr[i+1] {
					arr[i], arr[i+1] = arr[i+1], arr[i]
				}
			}
			if c == "desc" {
				if arr[i] < arr[i+1] {
					arr[i], arr[i+1] = arr[i+1], arr[i]
				}
			}

		}
		fmt.Println(arr)
	}
	return arr
}
