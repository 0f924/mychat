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
	fmt.Println("接收到单发信息响应包：", msgResp)
}
