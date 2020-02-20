package console

import (
	"fmt"
	"mychat/mychannel"
)

type SendToGroupConsoleCommand struct {
}

func (this SendToGroupConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	fmt.Println("发送群组消息")
}
