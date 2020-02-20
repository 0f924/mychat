package response

import (
	"mychat/mychannel"
	"mychat/protocol/packet"
)

type LoginResponseHandler struct {
}

func (this LoginResponseHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	loginResp := data.(packet.LoginResponsePacket)

	// 如果通过登录校验，记录登录状态
	if loginResp.IsSuccess {
		mychan.SetAttr("login", true)
	}
}
