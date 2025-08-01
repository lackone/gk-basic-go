package main

func myList(name string, aliases ...string) {
	println(name)

	for _, a := range aliases {
		println(a)
	}
}
