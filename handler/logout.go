package handler

import (
	"mychat/mychannel"
	"mychat/protocol/packet"
	"mychat/session"
)

type LogoutRequestHandler struct {
	Ctx *HandlerContext
}

func (this LogoutRequestHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	// 接收到退出登录请求包，即可删除服务端存储关于该客户端的信息
	user := mychan.GetAttr("user").(session.User)
	delete(this.Ctx.UserChan, user.UserId)

	// 回送响应包，通知客户端可以取消登录状态
	logoutResp := packet.LogoutResponsePacket{
		IsSuccess: true,
	}
	mychan.Write(logoutResp)
}

type LogoutResponseHandler struct {
}

func (this LogoutResponseHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	// 接收到退出登录响应包，即可取消客户端的登录状态
	mychan.DeleteAttr("login")
}
