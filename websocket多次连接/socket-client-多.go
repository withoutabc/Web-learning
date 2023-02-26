package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Fatalf("net.Dial err:%v", err)
		return
	}
	fmt.Println("client与server连接建立成功！")
	sendData := []byte("helloworld")

	for {
		//向服务器发送数据
		cnt, err := conn.Write(sendData)
		if err != nil {
			log.Fatalf("conn.Write err:%v", err)
			return
		}
		fmt.Println("Client ===> Server cnt:", cnt, ",data:", string(sendData))
		//接收服务器返回的数据
		//创建buf,用于接收服务器返回的数据
		buf := make([]byte, 1024)

		cnt, err = conn.Read(buf)
		if err != nil {
			log.Fatalf("conn.Read err:%v", err)
			return
		}
		fmt.Println("Client <==== Server,cnt:", cnt, ",data:", string(buf[:cnt]))
		time.Sleep(time.Second)
	}
	conn.Close()
}
