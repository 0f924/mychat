package handler

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
	"mychat/session"
)

// 请求包
type LoginRequestHandler struct {
	Ctx *HandlerContext
}

func (this LoginRequestHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	loginReq := data.(packet.LoginRequestPacket)

	// 通过登录校验，将用户信息保留到 Session
	user := session.User{
		UserId:   loginReq.UserId,
		UserName: loginReq.UserName,
	}
	mychan.SetAttr("user", user)

	this.Ctx.UserChan[user.UserId] = mychan
	for key, val := range this.Ctx.UserChan {
		fmt.Println("login:", key, "---", val)
	}

	// 回送：登录响应包
	loginResp := packet.LoginResponsePacket{
		IsSuccess: true,
		Reason:    "",
		UserId:    user.UserId,
		UserName:  user.UserName,
	}
	mychan.Write(loginResp)
}

// 响应包
type LoginResponseHandler struct {
}

func (this LoginResponseHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	loginResp := data.(packet.LoginResponsePacket)

	// 如果通过登录校验，记录登录状态
	if loginResp.IsSuccess {
		mychan.SetAttr("login", true)
	}
}
