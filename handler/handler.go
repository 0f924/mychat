package handler

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
)

type Handler interface {
	Exec(mychan *mychannel.MyChannel, data packet.Packet)
}

type HandlerManager struct {
	Ctx *HandlerContext
}

func (this HandlerManager) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	data = mychan.Read()
	switch data.GetType() {
	case packet.LOGIN_REQUEST:
		LoginRequestHandler{this.Ctx}.Exec(mychan, data)
	case packet.LOGIN_RESPONSE:
		LoginResponseHandler{}.Exec(mychan, data)
	case packet.MESSAGE_REQUEST:
		MessageRequestHandler{this.Ctx}.Exec(mychan, data)
	case packet.MESSAGE_RESPONSE:
		MessageResponseHandler{}.Exec(mychan, data)
	default:
		fmt.Println("识别不出数据包的类型！")
	}
}
