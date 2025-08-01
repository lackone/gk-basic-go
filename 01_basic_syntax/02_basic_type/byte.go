package main

import "fmt"

func Byte() {
	var a byte = 'a'
	println(a)

	fmt.Printf("%c\n", a)

	var s = "this is test"
	var sb = []byte(s)

	println(s)
	println(sb)
}
