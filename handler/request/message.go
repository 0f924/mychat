package request

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
	"mychat/session"
)

type MessageRequestHandler struct {
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

	// 多余代码
	fmt.Println(msgResp)

	// 发送响应包给目标用户

	// 待重构功能点
	// toChan := handler.UserIdChannelMap[msgReq.ToUserId]
	// toChan.Write(msgResp)
}
