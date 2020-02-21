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

	// 创建一个全局变量管理容器
	ctx := handler.NewHandlerContext()

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
		go pipeline(mychan, ctx)
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
func pipeline(mychan *mychannel.MyChannel, ctx *handler.HandlerContext) {
	startHandlePacket(mychan, ctx)
}

// 开启数据包监听处理服务
func startHandlePacket(mychan *mychannel.MyChannel, ctx *handler.HandlerContext) {
	defer mychan.Close()
	for {
		handler.HandlerManager{ctx}.Exec(mychan, nil)
	}
}
