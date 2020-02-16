package main

import (
	"fmt"
	"log"
	"mychat/protocol/channel"
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

	conn := connect(HOST, PORT, MAX_RETRY)

	handler := channel.ClientHandler{Conn: conn}
	handler.ChannelActive()
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		log.Fatal("客户端读取数据失败！")
	}
	handler.ChannelRead(data[:n])
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
