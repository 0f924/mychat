package main

import (
	"fmt"
	"log"
	"mychat/handler"
	"mychat/mychannel"
	"net"
	"time"
)

const (
	PORT string = "8081"
)

func main() {
	// 1. 监听端口
	listener := bind(PORT)
	defer listener.Close()

	// 2. 处理请求
	for {
		// 接收来自客户端的连接
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("与客户端的连接建立失败！")
		}

		// 装饰 服务端连接
		mychan := mychannel.NewMyChannel(conn)

		// 给每一个客户端分配一个服务端工作流水线
		pipeline(mychan)
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

// 服务端分配给每个连接的工作流水线
func pipeline(mychan *mychannel.MyChannel) {
	go startHandlePacket(mychan)
}

// 开启数据包监听处理服务
func startHandlePacket(mychan *mychannel.MyChannel) {
	handler.HandlerManager{}.Exec(mychan, nil)
}

// 接收 单发信息请求包
/*
func receiveMessageReq(conn net.Conn) {
	fmt.Println("等待用户传过来的信息.......")
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		log.Fatal("读取单发信息请求包失败！")
	}
	data = data[:n]
	req := packet.Decode(data).(packet.MessageRequestPacket)
	target_conn := session.UserIdChannelMap[req.ToUserId]
	source_user := session.ChannelUserMap[conn]
	resp := packet.MessageResponsePacket{
		FromUserId: source_user.UserId,
		FromUserName: source_user.UserName,
		Message: req.Message,
	}
	_, err = target_conn.Write(packet.Encode(resp))
	if err != nil {
		log.Fatal("单发信息响应包发送失败！")
	}
	fmt.Println("服务端已成功中转信息！")
}
*/
