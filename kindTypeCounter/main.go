package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("reader.Readline error = ", err)
		return
	}
	wc, sc, nc, oc := counter(string(line))
	fmt.Printf("word length: %d, space lenght: %d, number length: %d, others length: %d", wc, sc, nc, oc)
}

func counter(s string) (words, spaces, numbers, others int) {
	runes := []rune(s)

	for _, v := range runes {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			words++
		case v == ' ':
			spaces++
		case v >= '0' && v <= '9':
			numbers++
		default:
			others++
		}
	}
	return
}