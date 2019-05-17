package main

import (
	"fmt"
	"time"
)

func main() {
	//doFanIn1()
	//doFanIn2()
	doSelectChan()
}

func doFanIn1() {
	intChan := make(chan int)
	resultChan := make(chan int)
	exitChan := make(chan bool)
	// -> intChan
	go func() {
		for i := 0; i < 1000; i++ {
			intChan <- i
		}
		close(intChan)
	}()
	// intChan(1) -> resultChan(8) -> exitChan(1)
	for i := 0; i < 8; i++ {
		go calc(intChan, resultChan, exitChan)
	}
	// exitChan(1) ->
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resultChan)
	}()
	// resultChan(1) ->
	for v := range resultChan {
		fmt.Println("result: ", v)
	}
}

func calc(intChan, resultChan chan int, exitChan chan bool) {
	for v := range intChan {
		if v%3 == 0 {
			resultChan <- v
		}
	}
	exitChan <- true
}

func doFanIn2() {
	ch := make(chan int)
	exitChan := make(chan struct{})

	go send(ch, exitChan)
	go recv(ch, exitChan)

	var exitCount int
	for range exitChan {
		if exitCount == 1 {
			break
		}
		exitCount++
	}
}

func send(ch chan<- int, exitChan chan<- struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	exitChan <- struct{}{}
	fmt.Println("send exit.")
}

func recv(ch <-chan int, exitChan chan<- struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}

	exitChan <- struct{}{}
	fmt.Println("recv exit.")
}

func doSelectChan() {
	ch := make(chan int)
	ch2 := make(chan int)
	exit := make(chan bool)

	go func() {
		for i := 1; i < 10; i += 2 {
			ch <- i
		}
		exit <- true
	}()

	go func() {
		for i := 2; i < 10; i += 2 {
			ch2 <- i
		}
		time.Sleep(time.Second * 2)
		exit <- true
	}()

	var exitCount int
	for {
		select {
		case v := <-ch:
			fmt.Println("v:", v)
		case v2 := <-ch2:
			fmt.Println("v2:", v2)
		case <-exit:
			fmt.Println("exit.")
			exitCount++
			if exitCount == 2 {
				return
			}
		case <-time.After(time.Second):
			fmt.Println("timeout")
			return
		}
	}
}
