package console

import (
	"mychat/mychannel"
	"mychat/protocol/packet"
	"time"
)

type LogoutConsoleCommand struct {
}

func (this LogoutConsoleCommand) Exec(mychan *mychannel.MyChannel) {
	mychan.Write(packet.LogoutRequestPacket{})
	waitForLogout()
}

func waitForLogout() {
	time.Sleep(2 * time.Second)
}
