package request

import (
	"mychat/mychannel"
	"mychat/protocol/packet"
	"mychat/session"
	"mychat/utils"
)

type LoginRequestHandler struct {
}

func (this LoginRequestHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	loginReq := data.(packet.LoginRequestPacket)

	// 通过登录校验，将用户信息保留到 Session
	user := session.User{
		UserId:   loginReq.UserId,
		UserName: loginReq.UserName,
	}
	mychan.SetAttr("user", user)

	utils.Info("用户已登录")

	// 回送：登录响应包
	loginResp := packet.LoginResponsePacket{
		IsSuccess: true,
		Reason:    "",
		UserId:    user.UserId,
		UserName:  user.UserName,
	}
	mychan.Write(loginResp)
}
