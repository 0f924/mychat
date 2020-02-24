package mychannel

import (
	"mychat/protocol/packet"
	"mychat/utils"
	"net"
	"time"
)

const IDLE_TIME int = 15

type MyChannel struct {
	conn net.Conn
	attr map[string]interface{}
}

func NewMyChannel(conn net.Conn) *MyChannel {
	mychan := MyChannel{
		conn: conn,
		attr: make(map[string]interface{}),
	}
	mychan.UpdateDeadline()
	return &mychan
}

func (this *MyChannel) UpdateDeadline() {
	this.conn.SetDeadline(time.Now().Add(time.Duration(IDLE_TIME) * time.Second))
}

func (this *MyChannel) DeleteAttr(key string) {
	delete(this.attr, key)
}

func (this *MyChannel) SetAttr(key string, value interface{}) {
	this.attr[key] = value
}

func (this *MyChannel) GetAttr(key string) interface{} {
	return this.attr[key]
}

func (this *MyChannel) IsExists(key string) bool {
	if _, ok := this.attr[key]; !ok {
		return false
	}
	return true
}

func (this *MyChannel) Write(msg packet.Packet) {
	_, err := this.conn.Write(packet.Encode(msg))
	if err != nil {
		utils.Error("发送数据包失败！")
	}
}

func (this *MyChannel) Read() packet.Packet {
	buf := make([]byte, 1024)
	n, err := this.conn.Read(buf)
	if err != nil {
		utils.Error(err.Error())
	}
	buf = buf[:n]
	return packet.Decode(buf)
}

func (this *MyChannel) Close() {
	this.conn.Close()
	utils.Info("已关闭连接")
}
