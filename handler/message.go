package handler

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
	"mychat/session"
	"mychat/utils"
)

// 请求包
type MessageRequestHandler struct {
	Ctx *HandlerContext
}

func (this MessageRequestHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	msgReq := data.(packet.MessageRequestPacket)
	user := mychan.GetAttr("user").(session.User)
	// 构建：单发信息响应包
	msgResp := packet.MessageResponsePacket{
		FromUserId:   user.UserId,
		FromUserName: user.UserName,
		Message:      msgReq.Message,
	}

	for key, val := range this.Ctx.UserChan {
		fmt.Println("message:", key, "---", val)
	}

	// 发送响应包给目标用户
	if _, ok := this.Ctx.UserChan[msgReq.ToUserId]; !ok {
		utils.Info("用户:" + msgReq.ToUserId + "不存在")
		return
	}
	toChan := this.Ctx.UserChan[msgReq.ToUserId]
	toChan.Write(msgResp)
}

// 响应包
type MessageResponseHandler struct {
}

func (this MessageResponseHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	msgResp := data.(packet.MessageResponsePacket)
	fromUserId, fromUserName := msgResp.FromUserName, msgResp.FromUserName
	fmt.Println(fromUserName + "(" + fromUserId + "):" + msgResp.Message)
}
