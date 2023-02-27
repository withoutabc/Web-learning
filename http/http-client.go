package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//http包
	client := http.Client{}

	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		log.Fatalf("client.Get err:%v", err)
		return
	}

	body := resp.Body
	fmt.Println("body:", body)
	//func ReadAll(r io.Reader) ([]byte, error)
	readBodyStr, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatalf("read body err:%v", err)
		return
	}
	fmt.Println("body string:", string(readBodyStr))

	ct := resp.Header.Get("Content-Type")
	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")
	fmt.Println(ct)
	fmt.Println(date)
	fmt.Println(server)

	url := resp.Request.URL
	code := resp.StatusCode
	status := resp.Status
	fmt.Println(url)
	fmt.Println(code)
	fmt.Println(status)

}
