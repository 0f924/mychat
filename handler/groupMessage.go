package handler

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
	"mychat/session"
)

// 请求包
type GroupMessageRequestHandler struct {
	Ctx *HandlerContext
}

func (this GroupMessageRequestHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	grpMsgReq := data.(packet.GroupMessageRequestPacket)

	grpMsgResp := packet.GroupMessageResponsePacket{
		FromGroupId: grpMsgReq.ToGroupId,
		FromUser:    mychan.GetAttr("user").(session.User),
		Message:     grpMsgReq.Message,
	}

	chanGrp := this.Ctx.ChanGrp[grpMsgReq.ToGroupId]
	for _, userChan := range chanGrp {
		userChan.Write(grpMsgResp)
	}
}

// 响应包
type GroupMessageResponseHandler struct {
}

func (this GroupMessageResponseHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	grpMsgResp := data.(packet.GroupMessageResponsePacket)
	fromGroupId, fromUser := grpMsgResp.FromGroupId, grpMsgResp.FromUser
	fmt.Println("收到群[" + fromGroupId + "]中[" + fromUser.String() + "]发来的消息：" + grpMsgResp.Message)
}
