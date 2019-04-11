package main

import "fmt"

func main() {
	n := 5
	doWork(n)
}

func doWork(n int) {
	var sum int
	for i := 1; i <= n; i++ {
		sum += fact(i)
	}
	fmt.Println("sum = ", sum)
}

func fact(i int) int {
	switch i {
	case 0, 1:
		return 1
	default:
		return i * fact(i-1)
	}
}
