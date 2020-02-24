package main

import (
	"fmt"
	"mychat/client/console"
	"mychat/handler"
	"mychat/mychannel"
	"mychat/protocol/packet"
	"mychat/utils"
	"net"
	"strconv"
	"time"
)

const (
	MAX_RETRY int    = 5
	HOST      string = "127.0.0.1"
	PORT      string = "8081"
)

func main() {

	// 建立 tcp 连接
	conn := connect(HOST, PORT, MAX_RETRY)

	// 装饰 客户端连接
	mychan := mychannel.NewMyChannel(conn)

	// 启动 客户端工作流水线
	pipeline(mychan)
}

// 建立客户端连接，连接失败可重试
func connect(host, port string, retry int) net.Conn {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		if retry == 0 {
			panic("重试次数已用完，放弃连接！")
		} else {
			order := (MAX_RETRY - retry) + 1
			delay := order * 2
			fmt.Println(time.Now().String() + ": 连接失败，第" + strconv.Itoa(order) + "次重连......")
			time.Sleep(time.Second * time.Duration(delay))
			connect(host, port, retry-1)
		}
	}
	utils.Info("客户端连接成功")
	return conn
}

// 客户端工作流水线
func pipeline(mychan *mychannel.MyChannel) {
	go startHandlePacket(mychan)
	// 每 5 s发送一次心跳包
	go startHeartBeat(mychan, 5)
	startConsole(mychan)
}

// 开启用户命令行终端
func startConsole(mychan *mychannel.MyChannel) {
	loginCmd := console.LoginConsoleCommand{}
	loginedCmd := console.NewCommandManager()
	for {
		if mychan.IsExists("login") {
			loginedCmd.Exec(mychan)
		} else {
			loginCmd.Exec(mychan)
		}
	}
}

// 开启数据包监听处理服务
func startHandlePacket(mychan *mychannel.MyChannel) {
	defer mychan.Close()
	imhandler := handler.NewIMResponseHandler()
	for {
		imhandler.Exec(mychan, nil)
	}
}

// 开启心跳
func startHeartBeat(mychan *mychannel.MyChannel, idle int) {
	heartbeat := packet.HeartBeatRequestPacket{}
	for {
		mychan.Write(heartbeat)
		time.Sleep(time.Duration(idle) * time.Second)
	}
}
