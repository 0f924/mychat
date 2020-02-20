package console

import (
	"fmt"
	"mychat/mychannel"
)

type ListGroupMembersConsoleCommand struct {
}

func (this ListGroupMembersConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	fmt.Println("列出群组成员")
}
