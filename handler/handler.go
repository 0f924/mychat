package handler

import (
	"fmt"
	"mychat/handler/request"
	"mychat/handler/response"
	"mychat/mychannel"
	"mychat/protocol/packet"
)

type Handler interface {
	Exec(mychan *mychannel.MyChannel, data packet.Packet)
}

type HandlerManager struct {
}

func (this HandlerManager) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	data = mychan.Read()
	switch data.GetType() {
	case packet.LOGIN_REQUEST:
		request.LoginRequestHandler{}.Exec(mychan, data)
	case packet.LOGIN_RESPONSE:
		response.LoginResponseHandler{}.Exec(mychan, data)
	case packet.MESSAGE_REQUEST:
		request.MessageRequestHandler{}.Exec(mychan, data)
	case packet.MESSAGE_RESPONSE:
		response.MessageResponseHandler{}.Exec(mychan, data)
	default:
		fmt.Println("识别不出数据包的类型！")
	}
}
