package main

import (
	"fmt"
	"golang.org/x/net/context"
	"net"
	"net/http"
	"time"
)

func main() {
	url := []string{"http://www.baidu.com", "http://www.google.com", "http://taobao.com"}

	for _, v := range url {
		// set timeout control
		c := http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
					conn, e = net.DialTimeout(network, addr, time.Second*2)
					return
				},
			},
		}

		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err: %v\n", v, err)
			continue
		}
		fmt.Printf("get %s => successful status: %v\n", v, resp.Status)
	}
}
