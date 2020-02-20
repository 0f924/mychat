package console

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
	"mychat/utils"
	"time"
)

type LoginConsoleCommand struct {
}

func (this LoginConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	loginReq := packet.LoginRequestPacket{}

	fmt.Printf("输入用户名登录：")
	var username string
	fmt.Scanln(&username)
	loginReq.UserId = "wx" + utils.GetRandomId()
	loginReq.UserName = username
	loginReq.Password = "pwd"

	fmt.Println(username + ": " + loginReq.UserId)
	mychan.Write(loginReq)

	waitForLogin()
}

func waitForLogin() {
	time.Sleep(3 * time.Second)
}
