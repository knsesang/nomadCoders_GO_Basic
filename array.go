package main

import "fmt"

func main() {
	array_01()
	fmt.Println("------------")
	array_02()
	fmt.Println("------------")
}

func array_01() {
	names := [5]string{"aaa", "bbb", "ccc"}
	fmt.Println(names)
	// [aaa bbb ccc  ]

	names[3] = "ddd"
	names[4] = "eee"
	// names[5] = "fff" 오류 발생
	fmt.Println(names)
}

func array_02() {
	names := []string{"aaa", "bbb", "ccc"}
	fmt.Println(names)
	// [aaa bbb ccc  ]

	//  오류 발생
	// names[3] = "ddd"
	// names[4] = "eee"
	// names[5] = "fff"
	// fmt.Println(names)

	// 오류 발생
	// append(names, "ddd")

	names = append(names, "ddd")
	fmt.Println(names)
	// [aaa bbb ccc ddd]
}
