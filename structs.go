//  구조체
package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	struncts_01()
	fmt.Println("-------------")
	struncts_02()
	fmt.Println("-------------")
}

func struncts_01() {
	favFood := []string{"김치", "라면"}
	nico := person{"nico", 22, favFood}
	fmt.Println(nico)
	// {nico 22 [김치 라면]}
}

func struncts_02() {
	favFood := []string{"김치", "라면"}
	nico := person{name: "nico", age: 22, favFood: favFood}
	fmt.Println(nico)
	// {nico 22 [김치 라면]}
	fmt.Println(nico.name)
	// nico
}
