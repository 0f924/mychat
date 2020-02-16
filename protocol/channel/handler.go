package channel

import (
	"encoding/json"
	"fmt"
	"log"
	"mychat/protocol/packet"
	"mychat/utils"
	"net"
	"time"
)

type Handler interface {
	ChannelActive()
	ChannelRead(msg interface{})
}

// 登录客户端处理器
type ClientHandler struct {
	Conn net.Conn
}

func (this *ClientHandler) ChannelActive() {
	fmt.Println(time.Now().String() + ": 客户端开始登录")
	packet := packet.LoginRequestPacket{
		UserId:   utils.GetRandomId(),
		UserName: "flash",
		Password: "pwd",
	}
	buffer, err := json.Marshal(packet)
	if err != nil {
		log.Fatal("登录请求包编码失败！")
	}
	_, err = this.Conn.Write(buffer)
	if err != nil {
		log.Fatal("登录请求包发送失败！")
	}
}

func (this *ClientHandler) ChannelRead(data []byte) {
	var packet packet.LoginResponsePacket
	err := json.Unmarshal(data, &packet)
	if err != nil {
		log.Fatal("登录响应包解码失败！")
	}
	if packet.IsSuccess {
		fmt.Println(time.Now().String() + ": 客户端登录成功")
	} else {
		fmt.Println(time.Now().String() + ": 客户端登录失败，原因：" + packet.Reason)
	}
}

// 登录：服务端处理器
type ServerHandler struct {
	Conn net.Conn
}

func (this *ServerHandler) ChannelActive() {
}

func (this *ServerHandler) ChannelRead(data []byte) {
	fmt.Println(time.Now().String() + ": 客户端开始登录......")
	var req packet.LoginRequestPacket
	err := json.Unmarshal(data, &req)
	if err != nil {
		log.Fatal("登录请求包解码失败！")
	}
	fmt.Println(time.Now().String() + ": 登录成功！")
	resp := packet.LoginResponsePacket{
		IsSuccess: true,
		UserId:    req.UserId,
		UserName:  req.UserName,
	}
	resp_byte, err := json.Marshal(resp)
	if err != nil {
		log.Fatal("登录响应包编码失败！")
	}
	_, err = this.Conn.Write(resp_byte)
	if err != nil {
		log.Fatal("登录响应包发送失败！")
	}
}
