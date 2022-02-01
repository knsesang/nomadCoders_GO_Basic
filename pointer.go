package main

import "fmt"

func main() {
	pointer_01()
	fmt.Println("--------")
	pointer_02()
	fmt.Println("--------")
	pointer_03()
	fmt.Println("--------")
}

func pointer_01() {
	a := 2
	b := 3

	fmt.Println(a, b)
	// 2 3

	// 메모리 포인투 주소 출력
	fmt.Println(&a, &b)
	// 0xc000014098 0xc0000140b0

	fmt.Println(a + b)
}

func pointer_02() {
	a := 2
	b := &a

	fmt.Println(a, b)
	// 2 0xc0000140f0

	// 메모리 포인투 주소 출력
	fmt.Println(&a, &b)
	// 0xc000014098 0xc0000140b0
}

func pointer_03() {
	a := 2
	b := &a

	fmt.Println(a, &a, b, &b)
	// 2 0xc0000140f8 0xc0000140f8 0xc000006038

	*b = 20

	fmt.Println(a, &a, b, &b)
	// 20 0xc0000140f8 0xc0000140f8 0xc000006038
	// 메모리 값을 볼때는 *
	fmt.Println(a, &a, *b, &b)
	// 20 0xc0000140f8 20 0xc000006038
}
