package main

import (
	"fmt"
	"log"
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
	// 发送 登录请求
	sendLoginReq(conn)
	// 接收 登录响应
	receiveLoginResp(conn)
}

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
	fmt.Println(time.Now().String() + ": 连接成功！")
	return conn
}

// 发送 登录请求包
func sendLoginReq(conn net.Conn) {
	fmt.Println(time.Now().String() + ": 客户端开始登录")
	req := packet.LoginRequestPacket{
		UserId:   utils.GetRandomId(),
		UserName: "flash",
		Password: "pwd",
	}

	conn.Write(packet.Encode(req))
}

// 接收 登录响应包
func receiveLoginResp(conn net.Conn) {
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		log.Fatal("读取登录响应包失败！")
	}
	data = data[:n]
	resp := packet.Decode(data).(packet.LoginResponsePacket)
	if resp.IsSuccess {
		fmt.Println(time.Now().String() + ": 客户端登录成功")
	} else {
		fmt.Println(time.Now().String() + ": 客户端登录失败，原因：" + resp.Reason)
	}
}
