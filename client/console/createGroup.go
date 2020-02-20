package console

import (
	"fmt"
	"mychat/mychannel"
)

type CreateGroupConsoleCommand struct {
}

func (this CreateGroupConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	fmt.Println("创建群组")
}
