package main

func main() {
	swith_01()
	swith_02()
}

func swith_01() {
	age := 30
	switch age {

	case 10:
		println(10)
	case 20:
		println(20)
	case 30:
		println(30)
	}
}

func swith_02() {
	age := 30
	switch {

	case age == 10:
		println(10)
	case age == 20:
		println(20)
	case age == 30:
		println(30)
	}
}
