package main

import (
	"fmt"
	"log"
	"mychat/protocol/packet"
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
		go receiveLoginReq(conn)
	}
}

// 绑定端口：建立服务端服务
func bind(port string) net.Listener {
	listener, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		log.Fatal("端口[" + string(port) + "]绑定失败！")
	}
	fmt.Println(time.Now().String() + ": 端口[" + string(port) + "]绑定成功！")
	return listener
}

// 接收 登录请求包
func receiveLoginReq(conn net.Conn) {
	fmt.Println(time.Now().String() + ": 客户端开始登录......")
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		log.Fatal("读取登录请求包失败！")
	}
	data = data[:n]
	req := packet.Decode(data).(packet.LoginRequestPacket)
	fmt.Println(time.Now().String() + ": 登录成功！")
	resp := packet.LoginResponsePacket{
		IsSuccess: true,
		UserId:    req.UserId,
		UserName:  req.UserName,
	}
	_, err = conn.Write(packet.Encode(resp))
	if err != nil {
		log.Fatal("登录响应包发送失败！")
	}
}
