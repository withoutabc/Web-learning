package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//注册路由
	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	//func是回调函数，用于路由的响应
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		//request: ====> 包含客户端发来的数据
		fmt.Println(request)

		//writer: ===> 通过writer将数据返回给客户端
		_, _ = io.WriteString(writer, "这是/user请求返回的数据")
	})

	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {

	})

	http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {

	})
	fmt.Println("http server start ...")
	//func ListenAndServe(addr string, handler Handler) error
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		log.Fatalf("listen and serve err:%v", err)
		return
	}
}
