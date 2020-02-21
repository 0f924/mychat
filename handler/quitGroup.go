package handler

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
)

type QuitGroupRequestHandler struct {
	Ctx *HandlerContext
}

func (this QuitGroupRequestHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	quitGrpReq := data.(packet.QuitGroupRequestPacket)
	groupId := quitGrpReq.GroupId
	chanGrp := this.Ctx.ChanGrp[groupId]

	var idx int
	for idx := range chanGrp {
		if chanGrp[idx] == mychan {
			break
		}
	}
	chanGrp = append(chanGrp[:idx], chanGrp[idx+1:]...)
	this.Ctx.ChanGrp[groupId] = chanGrp

	quitGrpResp := packet.QuitGroupResponsePacket{
		IsSuccess: true,
		GroupId:   groupId,
	}
	mychan.Write(quitGrpResp)
}

type QuitGroupResponseHandler struct {
}

func (this QuitGroupResponseHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	quitGrpResp := data.(packet.QuitGroupResponsePacket)

	if quitGrpResp.IsSuccess {
		fmt.Println("退出群聊[" + quitGrpResp.GroupId + "]成功！")
	} else {
		fmt.Println("退出群聊[" + quitGrpResp.GroupId + "]失败！")
	}
}
