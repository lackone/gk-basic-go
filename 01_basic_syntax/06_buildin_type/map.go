package main

import "fmt"

func Map() {
	m := map[string]int{"foo": 1, "bar": 2}

	fmt.Println(m)

	m["foo"] = 3

	fmt.Println(m)

	m2 := make(map[string]int)
	m2["foo"] = 4
	m2["bar"] = 5
	fmt.Println(m2)

	if v, ok := m2["foo"]; ok {
		fmt.Println(v)
	}

}
