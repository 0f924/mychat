package console

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
	"strings"
)

const USER_ID_SPLITER string = ","

type CreateGroupConsoleCommand struct {
}

func (this CreateGroupConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	fmt.Printf("【拉人群聊】输入 userId 列表，userId 之间英文逗号隔开：")
	var userIds string
	fmt.Scanln(&userIds)

	crtGrpReq := packet.CreateGroupRequestPacket{
		UserIdList: strings.Split(userIds, USER_ID_SPLITER),
	}
	mychan.Write(crtGrpReq)
}
