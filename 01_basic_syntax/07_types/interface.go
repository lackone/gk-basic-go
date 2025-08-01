package main

type List interface {
	Len() int
	Add(index int, val any)
}

type LinkList struct {
}

type Integer int

func testInt() {
	i := 10
	i2 := Integer(i)
	println(i2)
}

func main() {

}
