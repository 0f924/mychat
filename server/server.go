package main

import (
	"fmt"
	"log"
	"mychat/protocol/channel"
	"net"
	"time"
)

const (
	PORT string = "8081"
)

func main() {
	// 1. 监听端口
	listener := bind(PORT)

	// 2. 处理请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("与客户端的连接建立失败！")
		}
		handler := channel.ServerHandler{Conn: conn}
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			log.Fatal("服务端读取数据失败！")
		}
		go handler.ChannelRead(data[:n])
	}
}

func bind(port string) net.Listener {
	listener, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		log.Fatal("端口[" + string(port) + "]绑定失败！")
	}
	fmt.Println(time.Now().String() + ": 端口[" + string(port) + "]绑定成功！")
	return listener
}
