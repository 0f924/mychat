package handler

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
)

// 请求包
type JoinGroupRequestHandler struct {
	Ctx *HandlerContext
}

func (this JoinGroupRequestHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	joinGrpReq := data.(packet.JoinGroupRequestPacket)

	groupId := joinGrpReq.GroupId
	chanGrp := this.Ctx.ChanGrp[groupId]
	chanGrp = append(chanGrp, mychan)
	this.Ctx.ChanGrp[groupId] = chanGrp

	joinGrpResp := packet.JoinGroupResponsePacket{
		IsSuccess: true,
		GroupId:   groupId,
	}
	mychan.Write(joinGrpResp)
}

// 响应包
type JoinGroupResponseHandler struct {
}

func (this JoinGroupResponseHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	joinGrpResp := data.(packet.JoinGroupResponsePacket)
	if joinGrpResp.IsSuccess {
		fmt.Println("加入群[" + joinGrpResp.GroupId + "]成功！")
	} else {
		fmt.Println("加入群[" + joinGrpResp.GroupId + "]失败，原因为：" + joinGrpResp.Reason)
	}
}
