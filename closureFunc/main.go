package main

import (
	"fmt"
	"strings"
)

func main() {
	f := addAccum()
	fmt.Println(f(1))
	fmt.Println(f(100))
	fmt.Println(f(1000))

	sf := makeSuffixFunc(".jpg")
	fmt.Println(sf("test"))
	fmt.Println(sf("test.jpg"))
}

func addAccum() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
