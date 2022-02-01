// func main() 을 우선적으로 찾게 된다
//	main.go 는 컴파일을 위한것이다
// Access is denied. 오류 발생시 코딩작업 폴더를 바이러스 검사 예외폴더로 지정한다
package main

// 10열에 fmt가 호출되면 자동으로 import 된다
import (
	"fmt"
)

func main() {
	fmt.Println("main start")

	account := banking.Account{}
	fmt.Println("main end")
}
