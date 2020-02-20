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

	var username string
	fmt.Printf("接收信息的用户名：")
	fmt.Scanln(&username)

	fmt.Printf("信息内容：")
	msg, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatal("读取用户输入的消息错误！")
	}

	msgReq := packet.MessageRequestPacket{
		ToUserId: username,
		Message:  msg,
	}

	mychan.Write(msgReq)

}
