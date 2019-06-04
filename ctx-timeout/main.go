package main

import (
	"fmt"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Result struct {
	res *http.Response
	err error
}

func main() {
	url := "http://www.imooc.com"
	proc(url)
}

func proc(url string) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	resChan := make(chan Result, 1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)

	go func() {
		resp, err := client.Do(req)
		pack := Result{res: resp, err: err}
		resChan <- pack
	}()

	select {
	case <-ctx.Done():
		res := <-resChan
		fmt.Printf("timeout: %v\n", res.err)
	case data := <-resChan:
		defer data.res.Body.Close()
		out, _ := ioutil.ReadAll(data.res.Body)
		fmt.Printf("data: %s", out)
	}
}
