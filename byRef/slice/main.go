package main

import "fmt"

func main() {
	list := []int{1, 2, 3}
	update(list)
	// todo: change the value by slice for slice.
	fmt.Println("main.list", list)
}

func update(list []int) {
	// slice by list
	list[0] = 2
	fmt.Println("set.list", list)
}
