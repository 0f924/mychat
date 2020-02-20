package console

import (
	"fmt"
	"mychat/mychannel"
)

type QuitGroupConsoleCommand struct {
}

func (this QuitGroupConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	fmt.Println("退出群组")
}
