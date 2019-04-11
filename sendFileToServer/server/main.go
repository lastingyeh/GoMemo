package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err = ", err)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	// todo#1 <- client [read filename]
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}
	fileName := string(buf[:n])

	// todo#2 -> client [send ok]
	conn.Write([]byte("ok"))

	// todo#3 <- client [received file content]
	err = RecvFile(fileName, conn)
	if err != nil {
		fmt.Println("RecvFile error = ", err)
		return
	}
}

func RecvFile(fileName string, conn net.Conn) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create error = ", err)
		return err
	}

	buf := make([]byte, 1024*4)
	for {
		n, err := conn.Read(buf)
		if err != nil && n != 0 {
			fmt.Println("conn.Read error = ", err)
			return err
		}

		if n == 0{
			fmt.Println("file received finished.")
			return nil
		}

		_, err = file.Write(buf[:n])
		if err != nil {
			fmt.Println("file.Write error = ", err)
			return err
		}
	}
}