// func main() 을 우선적으로 찾게 된다
//	main.go 는 컴파일을 위한것이다
// Access is denied. 오류 발생시 코딩작업 폴더를 바이러스 검사 예외폴더로 지정한다
package main

// 10열에 fmt가 호출되면 자동으로 import 된다
import (
	"fmt"
)

func main() {
	fmt.Println("-----------")
	repeatMe("a", "b", "c", "d")

	fmt.Println("for_ -----------")
	for_(1, 2, 3, 4, 5)
	fmt.Println("for_2 -----------")
	for_2(1, 2, 3, 4, 5)
	fmt.Println("for_3 -----------")
	for_3(1, 2, 3, 4, 5)

	fmt.Println("for_4 -----------")
	for_4(1, 2, 3, 4, 5)

	fmt.Println("for_5 -----------")
	for_5(1, 2, 3, 4, 5)

	fmt.Println("for_6 -----------")
	for_6(1, 2, 3, 4, 5)

	fmt.Println("for_7 -----------")
	for_7(1, 2, 3, 4, 5)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func for_(numbers ...int) {
	for number := range numbers {
		fmt.Println(number)
	}
}

func for_2(numbers ...int) {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
}

func for_3(numbers ...int) {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

func for_4(numbers ...int) {
	total := 0
	for number := range numbers {
		total += number
	}
	fmt.Println(total)
}

func for_5(numbers ...int) {
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println(total)
}

func for_6(numbers ...int) {
	for key, value := range numbers {
		fmt.Println(key, value)
	}
}

func for_7(numbers ...int) {
	for _, value := range numbers {
		fmt.Println(_, value)
	}
}
