package main

func Func1() {
	println("hello Func1")
}

func Func2() {
	fn := Func1
	fn()
}

// 函数变量
var aaa = func() {
	println("aaa")
}

// 返回一个函数
func Func3() func() {
	return func() {
		println("hello Func3")
	}
}
