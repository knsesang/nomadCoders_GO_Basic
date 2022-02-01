// func main() 을 우선적으로 찾게 된다
//	main.go 는 컴파일을 위한것이다
// Access is denied. 오류 발생시 코딩작업 폴더를 바이러스 검사 예외폴더로 지정한다
package main

// 10열에 fmt가 호출되면 자동으로 import 된다
import (
	"fmt"
)

func main() {
	if_01()
	if_02()
	if_03()
}

func if_01() {
	age := 23
	if age < 23 {
		fmt.Println("age11 ==========")
		fmt.Println(age)
	}
	fmt.Println("age11 ==========")
	fmt.Println(age)
}

func if_02() {
	if age2 := 23; age2 < 23 {
		fmt.Println("age21 ==========")
		fmt.Println(age2)
	}

	// if 문 밖에서 변수 호출 불가능
	// fmt.Println(age2)
}

func if_03() {
	x := 13
	if x > 10 {
		fmt.Println("++")
	} else if x == 10 {
		fmt.Println("10")
	} else {
		fmt.Println("--")
	}
}
