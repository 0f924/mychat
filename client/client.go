package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// 1. 打印“当前时间：客户端写出数据”，并发送“你好，闪电侠！”到服务端
	fmt.Println(time.Now().String() + ": 客户端写出数据")
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("你好，闪电侠！"))

	// 2. 读取服务端数据，并打印“当前时间：客户端读到数据 -> 数据”
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	result := buf[:n]
	fmt.Printf(time.Now().String() + ": 客户端读到数据 -> " + string(result))
}
