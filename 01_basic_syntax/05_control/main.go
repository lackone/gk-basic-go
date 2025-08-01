package main

func main() {
	age := 18

	if age >= 18 {
		println("成年")
	} else {
		println("未成年")
	}

	if age > 65 {
		println("老年")
	} else if age > 35 {
		println("中年")
	} else if age >= 18 {
		println("青年")
	}

	loop6()
}
