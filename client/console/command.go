package console

import (
	"fmt"
	"mychat/mychannel"
	"strings"
)

type Command interface {
	Exec(mychan *mychannel.MyChannel)
}

// 作用：分发任务
type CommandManager struct {
}

func (this CommandManager) Exec(mychan *mychannel.MyChannel) {
	fmt.Printf("请输入指令: ")
	var command string
	fmt.Scanln(&command)
	command = strings.ToLower(command)
	switch command {
	case "sendtouser":
		SendToUserConsoleCommand{}.Exec(mychan)
	case "sendtogroup":
		SendToGroupConsoleCommand{}.Exec(mychan)
	case "creategroup":
		CreateGroupConsoleCommand{}.Exec(mychan)
	case "joingroup":
		JoinGroupConsoleCommand{}.Exec(mychan)
	case "listgroupmembers":
		ListGroupMembersConsoleCommand{}.Exec(mychan)
	case "quitgroup":
		QuitGroupConsoleCommand{}.Exec(mychan)
	case "logout":
		LogoutConsoleCommand{}.Exec(mychan)
	default:
		fmt.Println("无法识别您输入的指令，请重新输入")
	}
}
