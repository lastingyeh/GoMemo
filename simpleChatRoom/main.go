package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string // chan client send message
	Name string      // client name
	Addr string      // addr name
}

var onlineMap map[string]*Client
var message = make(chan string)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}
	defer listener.Close()
	// collect msg, and send to all client
	go manageMessage()

	// main goroutine
	// todo: loop for client connect.
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	isQuit := make(chan bool)
	hasData := make(chan bool)
	// get client addr
	cliAddr := conn.RemoteAddr().String()
	// Client init
	cli := &Client{make(chan string), cliAddr, cliAddr}
	// add to map
	onlineMap[cliAddr] = cli
	// create goroutine 'client send message'
	go writeMsgToClient(cli, conn)
	// broadcast client online
	message <- makeMsg(cli, "login")
	cli.C <- makeMsg(cli, "I am here")

	// received client msg, create new goroutine
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				//client close || error
				isQuit <- true
				fmt.Println("conn.Read error = ", err)
				//return
			}
			// content send
			msg := buf[:n]
			s := strings.TrimSpace(string(msg))

			if len(s) == 3 && s == "who" {
				// send client individually.
				conn.Write([]byte("user list:\n"))

				for _, v := range onlineMap {
					list := fmt.Sprintf("%s:%s\n", v.Addr, v.Name)
					conn.Write([]byte(list))
				}
			} else if len(s) >= 8 && s[:6] == "rename" {
				name := strings.Split(s, "|")[1]
				cli.Name = name
				//onlineMap[cliAddr] = cli

				conn.Write([]byte("rename ok.\n"))
			} else {
				message <- makeMsg(cli, s)
			}

			hasData <- true
		}
	}()

	// lock handleConn
	for {
		select {
		case <-isQuit:
			s := makeMsg(cli, "leave chat room.")
			delete(onlineMap, cliAddr)
			message <- s
			return
		case <-hasData:
		case <-time.After(30 * time.Second):
			delete(onlineMap, cliAddr)
			message <- makeMsg(cli, "time out leave out.")
			return
		}
	}
}

func manageMessage() {
	// map init
	onlineMap = make(map[string]*Client)

	for {
		msg := <-message // lock as if no message
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func writeMsgToClient(cli *Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func makeMsg(cli *Client, msg string) string {
	return fmt.Sprintf("[%s]: %s", cli.Name, msg)
}