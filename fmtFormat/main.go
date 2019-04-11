package main

import "fmt"

type user struct {
	id   int
	name string
}

func main() {
	var a int
	var b bool
	fmt.Printf("a = %v, b = %v\n", a, b) // a = 0, b = false

	u := user{id: 1, name: "user"}
	fmt.Printf("%v\n", u)  // {1 user}
	fmt.Printf("%+v\n", u) // {id:1 name:user}
	fmt.Printf("%#v\n", u) // main.user{id:1, name:"user"}

	fmt.Printf("%d%%\n", 90)

	var c bool
	fmt.Printf("c = %t\n", c) // c = false

	var d string
	fmt.Printf("d = %q\n", d) // d = ""

	var e map[string]string
	var f []int
	fmt.Printf("e = %p\n", e) // e = 0x0
	fmt.Printf("f = %p\n", f) // f = 0x0
	//fmt.Printf("e == f: %t", e == f) // compile error

	str := fmt.Sprintf("%d", a)
	fmt.Printf("string(a) = %q", str) // string(a) = "0"
}