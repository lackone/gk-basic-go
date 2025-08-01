package main

//关键字func
//方法名大写，决定了作用域
//参数列表
//返回列表

//名字+参数列表+返回值，方法签名

func fn1() {

}

func fn2(a int) {

}

func fn3(a int, b int) {

}

// 有返回值，确保一定返回
func fn4() string {
	return ""
}

func fn5() (int, string) {
	return 1, "ok"
}

// 递归使用不当，会出现 stack overflow
func Recursive(n int) {
	if n > 10 {
		return
	}
	Recursive(n + 1)
}

func main() {
	//_ := fn4() //使用：=的前提，就是左边必须有至少一个新变量

	a, _ := fn5()
	println(a)

	Func2()

	aaa()

	fn := Func3()
	fn()

	closure := Closure("test")
	println(closure())

	Defer()
}
