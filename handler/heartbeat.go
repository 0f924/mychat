package handler

import (
	"mychat/mychannel"
	"mychat/protocol/packet"
)

type HeartBeatRequestHandler struct {
	Ctx *HandlerContext
}

func (this HeartBeatRequestHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	// 更新连接时间，然后返回响应包
	mychan.UpdateDeadline()
	if data.GetType() == packet.HEARTBEAT_REQUEST {
		mychan.Write(packet.HeartBeatResponsePacket{})
	}
}

type HeartBeatResponseHandler struct {
}

func (this HeartBeatResponseHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	// 更新连接时间
	mychan.UpdateDeadline()
}
