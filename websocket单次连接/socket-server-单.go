package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// 接受一个链接，而且只能发送一次数据
func main() {
	//创建监听
	ip := "127.0.0.1"
	port := 8000
	address := fmt.Sprintf("%s:%d", ip, port)
	//func Listen(network string, address string) (Listener, error)
	//net.Listener("tcp",":8080") //简写，冒号前默认是本机:127.0.0.1
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("net.listen err:%v", err)
		return
	}

	log.Println("监听中...")

	//func (Listener) Accept() (Conn, error)
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Listen.Accept err:%v", err)
		return
	}

	log.Println("建立连接成功！")

	//创建一个容器，用于接收读取到的数据
	buf := make([]byte, 1024) //使用make创建字节切片

	//func (Conn) Read(b []byte) (n int, err error)
	//cnt:真正读取client发来的数据长度
	cnt, err := conn.Read(buf)
	if err != nil {
		log.Fatalf("conn.Read err:%v", err)
		return
	}
	fmt.Println("Client ====> Server,长度：", cnt, ",data：", string(buf[:cnt]))
	//服务器对客户端请求进行响应，将数据转换成大写
	//func ToUpper(s string) string
	upperData := strings.ToUpper(string(buf[:cnt]))
	//func (Conn) Write(b []byte) (n int, err error)
	cnt, err = conn.Write([]byte(upperData))
	fmt.Println("Client ====> Server,长度：", cnt, ",data：", upperData)

	//关闭连接
	_ = conn.Close()
}
