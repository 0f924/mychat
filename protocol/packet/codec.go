package packet

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
)

const (
	MAGIC_NUMBER int64 = 0x12345678
	VERSION      int64 = 1
)

func Encode(packet Packet) []byte {
	// 1. JSON序列化数据对象
	data, err := json.Marshal(packet)
	if err != nil {
		log.Fatal("数据对象编码失败！")
	}

	// 2. 实际编码过程
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, MAGIC_NUMBER)
	binary.Write(buf, binary.BigEndian, VERSION)
	binary.Write(buf, binary.BigEndian, packet.GetType())
	binary.Write(buf, binary.BigEndian, rune(len(data)))
	binary.Write(buf, binary.BigEndian, data)

	return buf.Bytes()
}

func Decode(bys []byte) Packet {
	buf := bytes.NewBuffer(bys)

	// 跳过 magic number
	var magic_number int64
	err := binary.Read(buf, binary.BigEndian, &magic_number)
	if err != nil {
		log.Fatal("读取魔数失败！")
	}

	// 跳过 version
	var version int64
	err = binary.Read(buf, binary.BigEndian, &version)
	if err != nil {
		log.Fatal("读取版本号失败！")
	}

	// 获取 数据包类型
	var packetTypeCode byte
	err = binary.Read(buf, binary.BigEndian, &packetTypeCode)
	if err != nil {
		log.Fatal("读取数据包类型失败！")
	}

	// 读取 数据包长度
	var packetLength rune
	err = binary.Read(buf, binary.BigEndian, &packetLength)
	if err != nil {
		log.Fatal("读取数据包长度失败！")
	}

	// 获取 数据体
	data := make([]byte, int(packetLength))
	err = binary.Read(buf, binary.BigEndian, &data)
	if err != nil {
		log.Fatal("读取数据体失败！")
	}

	// JSON 解码
	var packet Packet
	switch packetTypeCode {
	case LOGIN_REQUEST:
		loginReq := LoginRequestPacket{}
		err = json.Unmarshal(data, &loginReq)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = loginReq

	case LOGIN_RESPONSE:
		loginResp := LoginResponsePacket{}
		err = json.Unmarshal(data, &loginResp)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = loginResp

	case LOGOUT_REQUEST:
		logoutReq := LogoutRequestPacket{}
		err = json.Unmarshal(data, &logoutReq)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = logoutReq

	case LOGOUT_RESPONSE:
		logoutResp := LogoutResponsePacket{}
		err = json.Unmarshal(data, &logoutResp)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = logoutResp

	case MESSAGE_REQUEST:
		msgReq := MessageRequestPacket{}
		err = json.Unmarshal(data, &msgReq)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = msgReq

	case MESSAGE_RESPONSE:
		msgResp := MessageResponsePacket{}
		err = json.Unmarshal(data, &msgResp)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = msgResp

	case CREATE_GROUP_REQUEST:
		crtGrpReq := CreateGroupRequestPacket{}
		err = json.Unmarshal(data, &crtGrpReq)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = crtGrpReq

	case CREATE_GROUP_RESPONSE:
		crtGrpResp := CreateGroupResponsePacket{}
		err = json.Unmarshal(data, &crtGrpResp)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = crtGrpResp

	case JOIN_GROUP_REQUEST:
		joinGrpReq := JoinGroupRequestPacket{}
		err = json.Unmarshal(data, &joinGrpReq)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = joinGrpReq

	case JOIN_GROUP_RESPONSE:
		joinGrpResp := JoinGroupResponsePacket{}
		err = json.Unmarshal(data, &joinGrpResp)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = joinGrpResp

	case LIST_GROUP_MEMBERS_REQUEST:
		lsGrpMebReq := ListGroupMembersRequestPacket{}
		err = json.Unmarshal(data, &lsGrpMebReq)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = lsGrpMebReq

	case LIST_GROUP_MEMBERS_RESPONSE:
		lsGrpMebResp := ListGroupMembersResponsePacket{}
		err = json.Unmarshal(data, &lsGrpMebResp)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = lsGrpMebResp

	case GROUP_MESSAGE_REQUEST:
		grpMsgReq := GroupMessageRequestPacket{}
		err = json.Unmarshal(data, &grpMsgReq)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = grpMsgReq

	case GROUP_MESSAGE_RESPONSE:
		grpMsgResp := GroupMessageResponsePacket{}
		err = json.Unmarshal(data, &grpMsgResp)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = grpMsgResp

	case QUIT_GROUP_REQUEST:
		quitGrpReq := QuitGroupRequestPacket{}
		err = json.Unmarshal(data, &quitGrpReq)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = quitGrpReq

	case QUIT_GROUP_RESPONSE:
		quitGrpResp := QuitGroupResponsePacket{}
		err = json.Unmarshal(data, &quitGrpResp)
		if err != nil {
			log.Fatal("JSON解码失败！")
		}
		packet = quitGrpResp
	}

	return packet
}
