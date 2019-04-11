package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

const url = "http://pengfu.163tc.com/xiaohua/index"
const listReRule = `<a title="#([^#]+)#" href="([^"]+)"`
const detailRule = `class="pic_text".*>[^>]*>(.*)</p>`

type Detail struct {
	title, content string
}

func main() {
	var start, end int
	fmt.Printf("start (>=1): ")
	fmt.Scan(&start)
	fmt.Printf("end (>= start page): ")
	fmt.Scan(&end)

	doWork(start, end)
}

func doWork(start, end int) {
	pageChan := make(chan int)
	fmt.Printf("crawl page from %d to %d\n", start, end)
	for i := start; i <= end; i++ {
		go spiderList(i, pageChan)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("list page %d loaded.\n", <-pageChan)
	}
}

func spiderList(i int, pageChan chan int) {
	var pageName string
	if i == 1 {
		pageName = ""
	} else {
		pageName = "_" + strconv.Itoa(i)
	}
	formatURL := fmt.Sprintf("%s%s.html", url, pageName)
	fmt.Printf("crawling %d page: %s\n", i, formatURL)

	result, err := httpGet(formatURL)
	if err != nil {
		fmt.Println("httpGet error = ", err)
		return
	}
	// todo: layer1 list
	list := contentParser(listReRule, result)
	details := make([]Detail, 0)
	detailsChan := make(chan []Detail)

	go func() {
		for _, item := range list {
			title, url := item[1], item[2]
			doc, err := httpGet(url)
			if err != nil {
				fmt.Println("detail.httpGet error = ", err)
				return
			}
			// todo: layer2 detail
			content := contentParser(detailRule, doc)
			//fmt.Println("content = ", content)
			if len(content) > 0 && len(content[0]) > 1 {
				detail := Detail{strings.TrimSpace(title), strings.TrimSpace(content[0][1])}
				details = append(details, detail)
				detailsChan <- details
			}
		}
		close(detailsChan)
	}()

	for data := range detailsChan {
		writeToFile(i, data)
	}
	pageChan <- i
}

func httpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get error = ", err)
		return "", err
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024*4)
	var bs bytes.Buffer
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		bs.Write(buf[:n])
	}
	return bs.String(), nil
}

func contentParser(rule, content string) [][]string {
	re := regexp.MustCompile(rule)
	matches := re.FindAllStringSubmatch(content, -1)
	return matches
}

func writeToFile(page int, details []Detail) {
	filename := path.Join("day08", "crawl", "pengfu", "src", strconv.Itoa(page)+".txt")
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create error = ", err)
		return
	}
	//var file *os.File
	//if _, err := os.Stat(filename); os.IsNotExist(err) {
	//	file, err = os.Create(filename)
	//	if err != nil {
	//		fmt.Println("os.Create error = ", err)
	//		return err
	//	}
	//} else {
	//	file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//	if err != nil {
	//		fmt.Println("os.OpenFile error = ", err)
	//		return err
	//	}
	//}
	defer file.Close()

	for _, detail := range details {
		file.WriteString(detail.title + "\n")
		file.WriteString(detail.content + "\n")
		file.WriteString("------------------------------------------------\n")
	}
}