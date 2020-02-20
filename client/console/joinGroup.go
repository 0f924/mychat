package console

import (
	"fmt"
	"mychat/mychannel"
)

type JoinGroupConsoleCommand struct {
}

func (this JoinGroupConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	fmt.Println("加入群组")
}
