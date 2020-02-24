package handler

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
)

type Handler interface {
	Exec(mychan *mychannel.MyChannel, data packet.Packet)
}

type IMRequestHandler struct {
	ctx        *HandlerContext
	handlerMap map[byte]Handler
}

func NewIMRequestHandler(ctx *HandlerContext) *IMRequestHandler {
	imHandler := &IMRequestHandler{
		ctx:        ctx,
		handlerMap: make(map[byte]Handler),
	}
	imHandler.handlerMap[packet.LOGIN_REQUEST] = LoginRequestHandler{ctx}
	imHandler.handlerMap[packet.LOGOUT_REQUEST] = LogoutRequestHandler{ctx}
	imHandler.handlerMap[packet.MESSAGE_REQUEST] = MessageRequestHandler{ctx}
	imHandler.handlerMap[packet.CREATE_GROUP_REQUEST] = CreateGroupRequestHandler{ctx}
	imHandler.handlerMap[packet.JOIN_GROUP_REQUEST] = JoinGroupRequestHandler{ctx}
	imHandler.handlerMap[packet.LIST_GROUP_MEMBERS_REQUEST] = ListGroupMembersRequestHandler{ctx}
	imHandler.handlerMap[packet.GROUP_MESSAGE_REQUEST] = GroupMessageRequestHandler{ctx}
	imHandler.handlerMap[packet.QUIT_GROUP_REQUEST] = QuitGroupRequestHandler{ctx}
	imHandler.handlerMap[packet.HEARTBEAT_REQUEST] = HeartBeatRequestHandler{ctx}

	return imHandler
}

func (this IMRequestHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	data = mychan.Read()
	packetType := data.GetType()

	if _, ok := this.handlerMap[packetType]; !ok {
		fmt.Println("识别不出数据包的类型！")
		return
	}
	if packetType != packet.HEARTBEAT_REQUEST {
		this.handlerMap[packet.HEARTBEAT_REQUEST].Exec(mychan, data)
	}
	this.handlerMap[packetType].Exec(mychan, data)
}

type IMResponseHandler struct {
	handlerMap map[byte]Handler
}

func NewIMResponseHandler() *IMResponseHandler {
	imHandler := &IMResponseHandler{
		handlerMap: make(map[byte]Handler),
	}
	imHandler.handlerMap[packet.LOGIN_RESPONSE] = LoginResponseHandler{}
	imHandler.handlerMap[packet.LOGOUT_RESPONSE] = LogoutResponseHandler{}
	imHandler.handlerMap[packet.MESSAGE_RESPONSE] = MessageResponseHandler{}
	imHandler.handlerMap[packet.CREATE_GROUP_RESPONSE] = CreateGroupResponseHandler{}
	imHandler.handlerMap[packet.JOIN_GROUP_RESPONSE] = JoinGroupResponseHandler{}
	imHandler.handlerMap[packet.LIST_GROUP_MEMBERS_RESPONSE] = ListGroupMembersResponseHandler{}
	imHandler.handlerMap[packet.GROUP_MESSAGE_RESPONSE] = GroupMessageResponseHandler{}
	imHandler.handlerMap[packet.QUIT_GROUP_RESPONSE] = QuitGroupResponseHandler{}
	imHandler.handlerMap[packet.HEARTBEAT_RESPONSE] = HeartBeatResponseHandler{}

	return imHandler
}

func (this IMResponseHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	data = mychan.Read()
	packetType := data.GetType()

	if _, ok := this.handlerMap[packetType]; !ok {
		fmt.Println("识别不出数据包的类型！")
		return
	}

	if packetType != packet.HEARTBEAT_RESPONSE {
		this.handlerMap[packet.HEARTBEAT_RESPONSE].Exec(mychan, data)
	}
	this.handlerMap[packetType].Exec(mychan, data)
}
