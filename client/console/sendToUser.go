package console

import (
	"bufio"
	"fmt"
	"log"
	"mychat/mychannel"
	"os"
)

type SendToUserConsoleCommand struct {
}

func (this SendToUserConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	inputReader := bufio.NewReader(os.Stdin)

	var username string
	fmt.Printf("接收信息的用户名：")
	fmt.Scanln(&username)
	fmt.Printf(username)

	fmt.Printf("信息内容：")
	msg, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatal("读取用户输入的消息错误！")
	}

	// 多余代码
	fmt.Println(msg)

	// 待重构功能点
	/**
	if _, ok := handler.NameToIdMap[username]; !ok {
		fmt.Println("不存在用户名" + username)
	}
	userId := handler.NameToIdMap[username]
	// fmt.Println("Send to user: ", userId)
	msgReq := packet.MessageRequestPacket{
		ToUserId: userId,
		Message:  msg,
	}

	mychan.Write(msgReq)

	*/
}
