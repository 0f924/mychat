package handler

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
	"mychat/session"
	"mychat/utils"
)

// 请求包
type CreateGroupRequestHandler struct {
	Ctx *HandlerContext
}

func (this CreateGroupRequestHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	crtGrpReq := data.(packet.CreateGroupRequestPacket)
	userIds := crtGrpReq.UserIdList

	// 根据群成员的id，获取其 mychannel
	chanGrp := make([]*mychannel.MyChannel, len(userIds))
	usernames := make([]string, len(userIds))
	for i, userId := range userIds {
		userChan := this.Ctx.UserChan[userId]
		chanGrp[i] = userChan
		user := userChan.GetAttr("user").(session.User)
		usernames[i] = user.UserName
	}

	// 回送群聊创建结果响应包
	groupId := utils.GetRandomId()
	crtGrpResp := packet.CreateGroupResponsePacket{
		IsSuccess:    true,
		GroupId:      groupId,
		UserNameList: usernames,
	}

	// 给每个群成员发拉群通知
	for _, userChan := range chanGrp {
		userChan.Write(crtGrpResp)
	}

	fmt.Printf("群创建成功，id 为 " + crtGrpResp.GroupId + ", ")
	fmt.Println("群里面有：", crtGrpResp.UserNameList)

	// 保存群组相关信息
	this.Ctx.ChanGrp[groupId] = chanGrp
}

// 响应包
type CreateGroupResponseHandler struct {
}

func (this CreateGroupResponseHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	crtGrpResp := data.(packet.CreateGroupResponsePacket)
	fmt.Print("群创建成功，id 为 [" + crtGrpResp.GroupId + "]，")
	fmt.Println("群里面有：", crtGrpResp.UserNameList)
}
