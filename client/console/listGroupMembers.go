package console

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
)

type ListGroupMembersConsoleCommand struct {
}

func (this ListGroupMembersConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	lsGrpMebReq := packet.ListGroupMembersRequestPacket{}

	fmt.Printf("输入 groupId，获取群成员列表：")
	var groupId string
	fmt.Scanln(&groupId)

	lsGrpMebReq.GroupId = groupId
	mychan.Write(lsGrpMebReq)
}
