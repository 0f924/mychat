package console

import (
	"bufio"
	"fmt"
	"log"
	"mychat/mychannel"
	"mychat/protocol/packet"
	"os"
)

type SendToGroupConsoleCommand struct {
}

func (this SendToGroupConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	fmt.Printf("发送消息给(群组ID)：")
	var toGroupId string
	fmt.Scanln(&toGroupId)

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("信息内容：")
	message, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatal("读取用户输入的消息错误！")
	}

	grpMsgReq := packet.GroupMessageRequestPacket{
		ToGroupId: toGroupId,
		Message:   message,
	}

	mychan.Write(grpMsgReq)
}
