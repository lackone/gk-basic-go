package main

// 闭包 = 函数 + 上下文
func Closure(name string) func() string {
	return func() string {
		return "hello " + name
	}
}
