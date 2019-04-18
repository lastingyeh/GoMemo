package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	for i := 1; i <= n; i++ {
		if isCompleteNum(i) {
			fmt.Printf("%d is complete Number\n", i)
		}
	}
}

func isCompleteNum(num int) bool {
	s := 0
	// 1. condition range 1 ~ n-1
	// 2. mod = 0
	// 3. sum(i) == n
	for i := 1; i < num; i++ {
		if num%i == 0 {
			s += i
		}
	}
	return num == s
}