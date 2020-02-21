package console

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
)

type QuitGroupConsoleCommand struct {
}

func (this QuitGroupConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	fmt.Printf("输入 groupId，退出群聊：")
	var groupId string
	fmt.Scanln(&groupId)

	quitGrpReq := packet.QuitGroupRequestPacket{
		GroupId: groupId,
	}
	mychan.Write(quitGrpReq)
}
