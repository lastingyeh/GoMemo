package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("reader.ReadLine error = ", err)
		return
	}

	// split +
	strSlice := strings.Split(string(line), "+")
	if len(strSlice) != 2 {
		fmt.Println("please input a + b")
		return
	}

	s1 := strings.TrimSpace(strSlice[0])
	s2 := strings.TrimSpace(strSlice[1])

	sum := add(s1, s2)
	fmt.Println("sum = ", sum)
}

func add(s1, s2 string) string {
	sLen := len(s1)
	if len(s1) > len(s2) {
		sLen = len(s2)
	}

	var addResult string
	var addCarry int
	var digits int

	rS1, rS2 := reverse(s1), reverse(s2)

	for i := 0; i < sLen; i++ {
		c1 := cInt(string(rS1[i]))
		c2 := cInt(string(rS2[i]))

		digits = (c1 + c2 + addCarry) % 10

		if c1+c2+addCarry > 9 {
			addCarry = 1
		} else {
			addCarry = 0
		}

		addResult += strconv.Itoa(digits)
	}

	resStr := extraStr(s1, s2, sLen, addCarry)
	addResult += resStr

	return reverse(addResult)
}

func extraStr(s1 string, s2 string, sLen int, addCarry int) string {
	switch {
	case len(s1) == len(s2):
		return strconv.Itoa(addCarry)
	default:
		s := s1
		var resStr string

		if len(s1) < len(s2) {
			s = s2
		}
		for i := 0; i < len(s)-sLen; i++ {
			resStr += string(s[i])
		}
		return reverse(strconv.Itoa(cInt(resStr) + addCarry))
	}
}

func cInt(b string) int {
	i, err := strconv.Atoi(b)
	if err != nil {
		panic(err)
	}
	return i
}

func reverse(s string) string {
	var bufs bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		bufs.WriteString(string(s[i]))
	}
	return bufs.String()
}