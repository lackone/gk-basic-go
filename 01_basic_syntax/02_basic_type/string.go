package main

import "unicode/utf8"

func String() {
	println("hello world")

	println(`可以换行
	我换行了
	我换行了
	我换行了`)

	println("hello" + "test")
	println("hello" + string(123)) //转成ASCII

	//len字节长度
	println(len("abc"))
	println(len("中国"))
	println(utf8.RuneCountInString("中国"))
}
