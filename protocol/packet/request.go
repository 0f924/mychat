package packet

// 请求数据包

// type: 创建群组
type CreateGroupRequestPacket struct {
	UserIdList []string
}

func (packet CreateGroupRequestPacket) GetType() byte {
	return CREATE_GROUP_REQUEST
}

// 创建群组 end

// type: 发送群组消息
type GroupMessageRequestPacket struct {
	ToGroupId string
	Message   string
}

func (packet GroupMessageRequestPacket) GetType() byte {
	return GROUP_MESSAGE_REQUEST
}

// 发送群组消息 end

// type: 心跳请求包
type HeartBeatRequestPacket struct {
}

func (packet HeartBeatRequestPacket) GetType() byte {
	return HEARTBEAT_REQUEST
}

// 心跳请求包 end

// type: 加入群组
type JoinGroupRequestPacket struct {
	GroupId string
}

func (packet JoinGroupRequestPacket) GetType() byte {
	return JOIN_GROUP_REQUEST
}

// 加入群组 end

// type: 列出群组成员
type ListGroupMembersRequestPacket struct {
	GroupId string
}

func (packet ListGroupMembersRequestPacket) GetType() byte {
	return LIST_GROUP_MEMBERS_REQUEST
}

// 列出群组成员 end

// type: 登录验证包
type LoginRequestPacket struct {
	UserId   string
	UserName string
	Password string
}

func (packet LoginRequestPacket) GetType() byte {
	return LOGIN_REQUEST
}

// 登录验证包 end

// type: 注销请求包
type LogoutRequestPacket struct {
}

func (packet LogoutRequestPacket) GetType() byte {
	return LOGOUT_REQUEST
}

// 注销请求包 end

// 单发消息
type MessageRequestPacket struct {
	ToUserId string
	Message  string
}

func (packet MessageRequestPacket) GetType() byte {
	return MESSAGE_REQUEST
}

// 单发消息 end

// type: 退出群组
type QuitGroupRequestPacket struct {
	GroupId string
}

func (packet QuitGroupRequestPacket) GetType() byte {
	return QUIT_GROUP_REQUEST
}

// 退出群组 end
