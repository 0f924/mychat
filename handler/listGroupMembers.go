package handler

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
	"mychat/session"
)

// 请求包
type ListGroupMembersRequestHandler struct {
	Ctx *HandlerContext
}

func (this ListGroupMembersRequestHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	lsGrpMebReq := data.(packet.ListGroupMembersRequestPacket)
	groupId := lsGrpMebReq.GroupId
	chanGrp := this.Ctx.ChanGrp[groupId]

	userList := make([]session.User, len(chanGrp))
	for i, userChan := range chanGrp {
		user := userChan.GetAttr("user").(session.User)
		userList[i] = user
	}

	lsGrpMebResp := packet.ListGroupMembersResponsePacket{
		GroupId:  groupId,
		UserList: userList,
	}
	mychan.Write(lsGrpMebResp)
}

// 响应包
type ListGroupMembersResponseHandler struct {
}

func (this ListGroupMembersResponseHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	lsGrpMebResp := data.(packet.ListGroupMembersResponsePacket)
	fmt.Println("群[" + lsGrpMebResp.GroupId + "]中的人包括：" + ToString(lsGrpMebResp.UserList))
}

func ToString(users []session.User) string {
	var result string = ""
	for _, user := range users {
		result = result + user.String() + ","
	}
	result = result[:len(result)-1]
	return result
}
