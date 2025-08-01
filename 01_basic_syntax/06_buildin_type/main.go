package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	fmt.Printf("%v len=%d cap=%d\n", arr1, len(arr1), cap(arr1))

	arr2 := [3]int{}
	fmt.Printf("%v len=%d cap=%d\n", arr2, len(arr2), cap(arr2))

	Slice()

	SubSlice()

	Map()
}
