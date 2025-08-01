package main

import "cmp"

func Sum[T cmp.Ordered](args ...T) T {
	var r T
	for _, arg := range args {
		r += arg
	}
	return r
}

type Integer int

type Num interface {
	int | float32 | float64 |
		~byte //~表示
}

func Sum2[T Num](args ...T) T {
	var r T
	for _, arg := range args {
		r += arg
	}
	return r
}

func main() {
	println(Sum(1, 2, 3))
	println(Sum2(1, 2, 3))
}
