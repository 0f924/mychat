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
	// 发送响应包给指定用户（待实现）
	fmt.Println("将要发送响应包:", msgResp)
	mychan.Write(msgResp)
}
