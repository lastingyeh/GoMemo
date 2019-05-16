package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

type charCount struct {
	char  int
	space int
	num   int
	other int
}

func main() {
	_charCounter()
	_readLine()
	_copy()
}

func _scanf() {
	var str = "jack 19 89.5"
	var stu Student

	_, err := fmt.Sscanf(str, "%s %d %f", &stu.Name, &stu.Age, &stu.Score)
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println(stu)
}

func _stdin() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("readString err = ", err)
	}
	fmt.Println(str)
}

func _charCounter() {
	file, err := os.Open("go_dev/ioFormat/main.go")
	if err != nil {
		fmt.Println("open file error = ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var count charCount
	for {
		str, err := reader.ReadString('\n')

		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Printf("readString error = %v", err)
			break
		}

		strArr := []rune(str)
		for _, v := range strArr {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.char++
			case v == ' ' || v == '\t':
				count.space++
			case v >= '0' && v <= '9':
				count.num++
			default:
				count.other++
			}
		}
	}
	fmt.Printf("%+v", count)
}

func _readLine() {
	file, err := os.Open("go_dev/ioFormat/main.go")
	if err != nil {
		fmt.Println("os.Open error: ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	index := 1
	var line []byte
	for {
		data, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		line = append(line, data...)
		if !isPrefix {
			fmt.Printf("line%d: %s\n", index, string(line))
			line = line[:0]
		}
		index++
	}
}

func _copy() {
	src, err := os.Open("go_dev/ioFormat/main.go")
	if err != nil {
		fmt.Println("src open error: ", err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile("go_dev/readline/copy.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("dst open error: ", err)
		return
	}
	defer dst.Close()
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println("copy error: ", err)
	}

}
