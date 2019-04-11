package main

import "fmt"

func main() {
	a := 5
	b := make(chan int, 1)

	fmt.Println("a = ", a) // value
	fmt.Println("b = ", b) // ref

	updateValue(1)
	fmt.Println("update value for a = ", a) // unchanged

	updateRef(&a)
	fmt.Println("update ref by for a = ", a) // changed
}

func updateValue(a int) {
	a = 10
}

func updateRef(a *int) {
	*a = 10
}
