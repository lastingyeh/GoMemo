package main

import "fmt"

func main() {
	ch := make(chan int)    // int chan
	quit := make(chan bool) // quit flag

	// consumer int <-chan
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println("num = ", num)
		}
		quit <- true
	}()
	// producer chan<- int
	fib(ch, quit)
}

func fib(ch chan<- int, quit chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}
}