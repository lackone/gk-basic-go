package main

import "math"

func main() {
	var a = 123
	var b = 456
	println(a + b)
	println(a - b)
	println(a * b)
	println(a / b)

	var c = 3.14
	println(a + int(c)) //无法自动类型转换

	println(math.MaxInt32)
	println(math.MinInt32)

	String()

	Byte()
}
