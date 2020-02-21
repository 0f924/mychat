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
		fmt.Println("无法识别您输入的指令，请输以下指令：")
		fmt.Println("sendToUser --- 发送信息给用户")
		fmt.Println("createGroup --- 创建群聊")
		fmt.Println("joinGroup --- 加入群聊")
		fmt.Println("sendToGroup --- 发送群聊信息")
		fmt.Println("listGroupMembers --- 列出群聊成员")
		fmt.Println("quitGroup --- 退出群聊")
		fmt.Println("logout --- 退出登录")
	}
}
