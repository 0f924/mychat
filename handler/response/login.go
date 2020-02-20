package response

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
	"mychat/utils"
)

type LoginResponseHandler struct {
}

func (this LoginResponseHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	loginResp := data.(packet.LoginResponsePacket)
	fmt.Println("接收到登录响应包：", loginResp)

	// 如果通过登录校验，记录登录状态
	if loginResp.IsSuccess {
		mychan.SetAttr("login", true)
		utils.Info("用户登录成功")
	}
}
