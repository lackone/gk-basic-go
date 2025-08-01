package main

func SwitchAge(age int) {
	switch {
	case age >= 18:
		println("青年")
	case age >= 65:
		println("老年")
	case age >= 35:
		println("中年")
	}
}
