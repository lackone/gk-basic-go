package main

func Defer() {
	defer func() {
		println("aaa")
	}()

	defer func() {
		println("bbb")
	}()
}

func Defer2() {
	i := 0
	defer func() {
		println(i)
	}()
	i = 1
}

func Defer3() {
	i := 0
	defer func(v int) {
		println(v)
	}(i)
	i = 1
}

func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturn2() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}
