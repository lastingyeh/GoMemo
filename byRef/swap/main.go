package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Printf("swap: a = %p, b = %p\n", &a, &b)
	swap(&a, &b)

	fmt.Printf("main: a = %d, b = %d\n", a, b)
}

func swap(a, b *int) {
	fmt.Printf("swap: a = %p, b = %p\n", a, b)
	fmt.Printf("swap: &a = %p, &b = %p\n", &a, &b)
	*a, *b = *b, *a
	fmt.Printf("swap: a = %d, b = %d\n", *a, *b)
}
