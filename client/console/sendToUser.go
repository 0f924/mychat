package console

import (
	"bufio"
	"fmt"
	"log"
	"mychat/mychannel"
	"mychat/protocol/packet"
	"os"
)

type SendToUserConsoleCommand struct {
}

func (this SendToUserConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Printf("发送消息给flash：")

	msg, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatal("读取用户输入的消息错误！")
	}

	msgReq := packet.MessageRequestPacket{
		ToUserId: "flash",
		Message:  msg,
	}
	mychan.Write(msgReq)
}
