package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// set seed to random number
	rand.Seed(time.Now().Unix())
	n := rand.Intn(100)
	fmt.Println(n)

	for {
		var input int
		fmt.Scanf("%d\n", &input) // \n <- omit return key scan.
		flag := false

		switch {
		case n > input:
			fmt.Printf("bigger than %d\n", input)
		case n < input:
			fmt.Printf("smaller than %d\n", input)
		case n == input:
			fmt.Printf("correct: %d\n", input)
			flag = true
			//break <- can't break for-loop
		}

		if flag {
			break
		}
	}
}