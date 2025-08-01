package main

import "fmt"

func Slice() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	s2 := make([]int, 3, 5)
	fmt.Println(s2)
	s2 = append(s2, 14)
	fmt.Println(s2)

	s3 := make([]int, 0, 5)
	fmt.Println(s3)
	s3 = append(s3, 15)
	fmt.Println(s3)
}

// 左闭右开原则
func SubSlice() {
	s := []int{2, 3, 5, 7, 11, 13}

	s2 := s[1:4]
	fmt.Println(s2)

	s3 := s[2:]
	fmt.Println(s3)

	s4 := s[:2]
	fmt.Println(s4)
}
