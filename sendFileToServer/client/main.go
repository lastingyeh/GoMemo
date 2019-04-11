package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// show input filename hint.
	fmt.Println("input file:")
	var path string
	fmt.Scan(&path)
	// get filename
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err = ", err)
		return
	}

	// connect to server
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("conn to server err = ", err)
		return
	}
	defer conn.Close()

	// todo#1 -> server [send filename first]
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("conn.Write error =", err)
		return
	}

	// todo#2 <- server [received from server, confirm ok?]
	var n int
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}

	if string(buf[:n]) == "ok" {
		//todo#3 -> server [sendfile to server]
		err = SendFile(path, conn)
		if err != nil {
			fmt.Println("Sendfile error = ", err)
		}
	}
}

func SendFile(path string, conn net.Conn) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	// read file content
	buf := make([]byte, 1024*4)
	for {
		n, err := file.Read(buf)

		if err == io.EOF {
			fmt.Println("read all.")
			return nil
		}

		if err != nil {
			fmt.Println("file.Read err = ", err)
			return err
		}

		_, err = conn.Write(buf[:n])
		if err != nil {
			return err
		}
	}
}