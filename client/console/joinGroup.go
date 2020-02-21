package console

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
)

type JoinGroupConsoleCommand struct {
}

func (this JoinGroupConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	joinGrpReq := packet.JoinGroupRequestPacket{}

	fmt.Printf("输入 groupId，加入群聊：")
	var groupId string
	fmt.Scanln(&groupId)

	joinGrpReq.GroupId = groupId
	mychan.Write(joinGrpReq)
}
