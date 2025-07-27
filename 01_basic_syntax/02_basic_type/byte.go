package main

import "fmt"

func Byte() {
	var a byte = 'a'
	println(a)

	fmt.Printf("%c\n", a)

	var s string = "this is test"
	var sb []byte = []byte(s)

	println(s)
	println(sb)
}
