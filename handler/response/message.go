package response

import (
	"fmt"
	"mychat/mychannel"
	"mychat/protocol/packet"
)

type MessageResponseHandler struct {
}

func (this MessageResponseHandler) Exec(mychan *mychannel.MyChannel, data packet.Packet) {
	msgResp := data.(packet.MessageResponsePacket)
	fromUserId, fromUserName := msgResp.FromUserName, msgResp.FromUserName
	fmt.Println(fromUserId + ":" + fromUserName + " -> " + msgResp.Message)
}
