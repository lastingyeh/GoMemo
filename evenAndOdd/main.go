package main

import "fmt"

func main() {
	var start, end int
	var nType string
	fmt.Scanf("%s%d%d", &nType, &start, &end) // args input: 'odd 1 2'
	fmt.Println(nType, start, end)

	doWork(nType, start, end)
}

func doWork(nType string, start, end int) {
	fmt.Println(nType)
	var mod int

	switch nType {
	case "odd":
		mod = 1
	case "even":
		mod = 0
	default:
		panic("nType not found, it must enum 'odd' or 'even'")
	}

	for i := start; i <= end; i++ {
		if i%2 == mod {
			fmt.Println("num = ", i)
		}
	}
}