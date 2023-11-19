package main

import "fmt"

// Slice 注意和 Array 的区别，就是[] 里不填数值，就是切片了
// 注意：len 为已经放了多少，cap 是可以放多少
// 因此，在初始化切片时，要预估容量
func Slice() {
	s1 := []int{9, 8, 7}
	fmt.Printf("s1: %v, len=%d, cap=%d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 4)
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))

	s3 := make([]int, 2)
	fmt.Printf("s3: %v, len=%d, cap=%d \n", s3, len(s3), cap(s3))

	s4 := make([]int, 0, 4)
	s4 = append(s4, 1)
	fmt.Printf("s4: %v, len=%d, cap=%d \n", s4, len(s4), cap(s4))

	// s1: [9 8 7], len=3, cap=3
	// s2: [0 0 0], len=3, cap=4
	// s3: [0 0 0 0], len=4, cap=4
	// s4: [1], len=1, cap=4
}

func SubSlice() {
	s1 := []int{2, 4, 6, 8, 10}
	s2 := s1[1:3] // 左闭右开
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))
	// s2: [4 6], len=2, cap=4

	// 只有 start  或者只有 end
	s3 := s1[0:]
	fmt.Printf("s3: %v, len=%d, cap=%d \n", s3, len(s3), cap(s3))
	s4 := s1[:5]
	fmt.Printf("s4: %v, len=%d, cap=%d \n", s4, len(s4), cap(s4))
	// s3: [2 4 6 8 10], len=5, cap=5
	// s4: [2 4 6 8 10], len=5, cap=5
}

// ShareSlice 切片和子切片，共享内存
// 扩容后，就不在共享内存
func ShareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))

	s2[0] = 99
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))
	fmt.Printf("s1: %v, len=%d, cap=%d \n", s1, len(s1), cap(s1))
	//s2: [3 4], len=2, cap=2
	//s2: [99 4], len=2, cap=2
	//s1: [1 2 99 4], len=4, cap=4

	s2 = append(s2, 199)
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))
	fmt.Printf("s1: %v, len=%d, cap=%d \n", s1, len(s1), cap(s1))
	// s2: [3 4], len=2, cap=2
	// s2: [99 4], len=2, cap=2
	// s1: [1 2 99 4], len=4, cap=4
	// s2: [99 4 199], len=3, cap=4
	// s1: [1 2 99 4], len=4, cap=4

	s2[1] = 1999
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))
	fmt.Printf("s1: %v, len=%d, cap=%d \n", s1, len(s1), cap(s1))
	// s2: [3 4], len=2, cap=2
	// s2: [99 4], len=2, cap=2
	// s1: [1 2 99 4], len=4, cap=4
	// s2: [99 4 199], len=3, cap=4
	// s1: [1 2 99 4], len=4, cap=4
	// s2: [99 1999 199], len=3, cap=4
	// s1: [1 2 99 4], len=4, cap=4
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	firstValue := first(nums, target)
	if firstValue == -1 {
		return []int{-1, -1}
	}
	lastValue := last(nums, target)
	return []int{firstValue, lastValue}

}

func last(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left+1)/2
		if nums[mid] == target {
			left = mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return left
}

func first(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}
