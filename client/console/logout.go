package console

import (
	"fmt"
	"mychat/mychannel"
)

type LogoutConsoleCommand struct {
}

func (this LogoutConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	fmt.Println("用户退出登录")
}
