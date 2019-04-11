package main

import (
	"fmt"
)

func main() {
	start, end := 100, 999
	doWork(start, end)
}

func doWork(start, end int) {
	for i := start; i <= end; i++ {
		div, mod := i, 0
		var nSum int
		for {
			div, mod = div/10, div%10
			nSum += mod * mod * mod

			if div == 0 {
				break
			}
		}
		//fmt.Printf("nSum = %d, i = %d\n", nSum, i)

		// check 
		if nSum == i {
			fmt.Printf("getNarcissisticNumber: %d\n", i)
		}
	}
}