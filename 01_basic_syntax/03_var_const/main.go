package main

// 首字符大小写控制访问性
var Global = "全局变量，外部可访问"
var internal = "全局变量，外部不可访问"

const PI = 3.1415926
const pi = 3.1415926

const (
	test string = "test"
	code        = "500"
)

const (
	statusA = iota
	statusB
	statusC
	statusD = 999
	statusE
)

func main() {
	var a int = 111
	println(a)
	var b = 222
	println(b)
	c := 333
	println(c)

	var (
		d int
		e string
		f float64
	)
	println(d, e, f)

	println(statusA, statusB, statusC, statusD, statusE)
}
