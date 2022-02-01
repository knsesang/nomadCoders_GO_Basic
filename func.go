package main

import (
	"fmt"
	"strings"
)

// # 방법 1
// func lenAndUpper(name string) (int, string) {
//     return len(name), strings.ToUpper(name)
// }

// #방법 2
// func lenAndUpper(name string) int {
//     return len(name)
// }

// # 방법 3
// func lenAndUpper2(name string) (length int, uppercase string) {
//     length = len(name)
//     uppercase = strings.ToUpper(name)
//     return
// }

func main() {
	fmt.Println("01 ---------------")
	totalLenght, upperName := func_01("nico")
	fmt.Println(totalLenght, upperName)

	fmt.Println("02 ---------------")
	totalLenght2, upperName2 := func_02("nico")
	fmt.Println(totalLenght2, upperName2)

	fmt.Println("03 ---------------")
	fmt.Println("main start")
	func_03("nico")
	fmt.Println("main end")
}

func func_01(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func func_02(name string) (length int, uppercase string) {
	return len(name), strings.ToUpper(name)
}

func func_03(name string) (length int, uppercase string) {
	defer fmt.Println("defer print")
	length = len(name)
	uppercase = strings.ToUpper(name)
	fmt.Println("func run complete")
	return
}
