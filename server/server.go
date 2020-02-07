package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// 1. 接收从客户端获取的数据，并打印“当前时间：服务端读到数据 -> 数据”
	listenner, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer listenner.Close()
	conn, err := listenner.Accept()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	result := buf[:n]
	fmt.Println(time.Now().String() + ": 服务端读到数据 -> " + string(result))

	// 2. 打印“当前时间：服务端写出数据”，并发送“你好，欢迎关注我的微信公众号，《闪电侠的博客》！”
	fmt.Println(time.Now().String() + ": 服务端写出数据")
	conn.Write([]byte("你好，欢迎关注我的微信公众号，《闪电侠的博客》！"))
}
