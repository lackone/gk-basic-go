package main

import "fmt"

func loop1() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}

func loop2() {
	i := 0
	for i < 10 {
		println(i)
		i++
	}
}

func loop3() {
	for {
		println("loop3")
	}
}

func loop4() {
	a := []int{1, 2, 3}

	for i, v := range a {
		println(i, v)
	}
}

func loop5() {
	m := map[string]int{"one": 1, "two": 2, "three": 3}

	for k, v := range m {
		println(k, v)
	}
}

type user struct {
	name string
}

func loop6() {
	users := []user{
		{name: "aaaa"},
		{name: "bbbb"},
	}

	m := make(map[string]*user)

	for _, v := range users {
		m[v.name] = &v
	}

	fmt.Printf("%v\n", m)
}
