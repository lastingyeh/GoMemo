package main

import "fmt"

type opFunc func(string, ...int)

func main() {
	operator(add, "add", 200, 300)
	operator(sub, "sub", 500, 200)
	operator(mul, "mul", 50, 40)
}

// same addFunc structure
func add(funcName string, args ...int) {
	var res int

	for _, arg := range args {
		res += arg
	}

	fmt.Printf("func= %s, result = %d\n", funcName, res)
}

//  same addFunc structure
func sub(funcName string, args ...int) {
	res := args[0]

	for idx, arg := range args {
		if idx == 0 {
			continue
		}
		res -= arg
	}

	fmt.Printf("func= %s, result = %d\n", funcName, res)
}

//  same addFunc structure
func mul(funcName string, args ...int) {
	res := args[0]

	for idx, arg := range args {
		if idx == 0 {
			continue
		}
		res *= arg
	}

	fmt.Printf("func= %s, result = %d\n", funcName, res)
}

func operator(op opFunc, name string, args ...int) {
	op(name, args...)
}